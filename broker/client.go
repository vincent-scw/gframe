package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	k "github.com/vincent-scw/gframe/kafka"
)

func main() {
	log.Println("Starting kafka broker...")

	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	handler := newReceptionHandler()
	consumer := k.Consumer{
		Ready:   make(chan bool, 0),
		Handler: handler,
	}

	ctx, cancel := context.WithCancel(context.Background())

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup([]string{"40.83.112.48:9092"}, "player_broker", config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{"player"}, &consumer); err != nil {
				log.Panicf("Error form consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.Ready = make(chan bool, 0)
		}
	}()

	<-consumer.Ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

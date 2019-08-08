package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	k "github.com/vincent-scw/gframe/kafkactl"
	"github.com/vincent-scw/gframe/util"
)

func main() {
	log.Println("Starting broker service...")

	// Set default configurations
	viper.SetDefault("REDIS_SERVER", "localhost:6379")
	viper.SetDefault("KAFKA_BROKERS", []string{"localhost:9092"})

	viper.AutomaticEnv() // automatically bind env

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

	var client sarama.ConsumerGroup
	err = util.WithRetry(10, 2*time.Second, func() (err error) {
		client, err = sarama.NewConsumerGroup(viper.GetStringSlice("KAFKA_BROKERS"), "player_broker", config)
		return
	})
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

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/vincent-scw/gframe/broker_svc/config"
	k "github.com/vincent-scw/gframe/kafkactl"
	"github.com/vincent-scw/gframe/util"
)

func main() {
	log.Println("Starting broker service...")

	go serveWeb(config.GetWebPort())

	ctx, cancel := context.WithCancel(context.Background())

	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}
	handler := newReceptionHandler()
	consumer := k.Consumer{
		Ready:   make(chan bool, 0),
		Handler: handler,
	}
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

	//serveRPC(config.GetRPCPort())
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

// func serveRPC(port string) {
// 	handler := newReceptionHandler()
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	e.RegisterUserReceptionServer(s, handler)
// 	log.Printf("Listen to RPC at port %s", port)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func serveWeb(port string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am healthy.")
	})
	log.Println(fmt.Sprintf("Serve start at %s.", port))
	log.Fatal(http.ListenAndServe(port, nil))
}

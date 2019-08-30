package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/vincent-scw/gframe/broker_svc/config"
	k "github.com/vincent-scw/gframe/kafkactl"
	"github.com/vincent-scw/gframe/util"
)

func main() {
	log.Println("Starting broker service...")

	ctx, cancel := context.WithCancel(context.Background())

	consumer := k.Consumer{
		Ready:   make(chan bool, 0),
		Handler: newReceptionHandler(),
		Context: ctx,
	}

	var client *k.ConsumerGroup
	err := util.WithRetry(10, 2*time.Second, func() (err error) {
		client, err = k.NewConsumerGroup(config.GetKafkaBrokers(), k.TopicPlayer, &consumer)
		return
	})
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	defer client.Close()

	log.Println("Sarama consumer up and running...")

	serveWeb(config.GetWebPort())

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

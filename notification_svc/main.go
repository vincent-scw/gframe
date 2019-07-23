package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

func main() {
	log.Println("Starting player notification service...")

	pubsub := r.NewPubSubClient("40.83.112.48:6379")
	defer pubsub.Close()

	log.Println("Subscribing to Redis...")
	pubsub.Subscribe(e.GroupChannel, func(msg string) {

	})

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	log.Println("terminating: via signal")
}

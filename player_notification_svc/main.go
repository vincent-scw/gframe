package main

import (
	"log"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

func main() {
	log.Println("Starting player notification service...")

	pubsub := r.NewPubSubClient("40.83.112.48:6379")
	defer pubsub.Close()

	pubsub.Subscribe(e.GroupChannel, func(msg string) {
		log.Println(msg)
	})
}

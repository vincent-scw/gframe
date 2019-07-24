package main

import (
	"log"
	"net/http"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

func main() {
	log.Println("Starting player notification service...")

	pubsub := r.NewPubSubClient("40.83.112.48:6379")
	defer pubsub.Close()

	hub := newHub()
	go hub.run()

	log.Println("Subscribe to Redis...")
	go pubsub.Subscribe(e.GroupChannel, func(msg string) {
		hub.broadcast <- []byte(msg)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am good."))
	})

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Println("Serve at localhost:9010...")
	log.Fatal(http.ListenAndServe(":9010", nil))
}

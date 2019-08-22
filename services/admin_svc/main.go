package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"

	r "github.com/vincent-scw/gframe/redisctl"

	"github.com/vincent-scw/gframe/admin_svc/config"
	gql "github.com/vincent-scw/gframe/admin_svc/graphql"
	"github.com/vincent-scw/gframe/admin_svc/subscriber"
)

func main() {
	log.Println("Starting admin service...")

	pubsub := r.NewPubSubClient(config.GetRedisServer())
	defer pubsub.Close()

	log.Println("Subscribe to Redis...")
	go subscriber.SubscribePlayer(pubsub, gql.Broadcast)
	go subscriber.SubscribeGroup(pubsub, gql.Broadcast)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am good."))
	})

	h := handler.New(&handler.Config{
		Schema: &gql.Schema,
		Pretty: true,
	})
	mux.Handle("/graphql", h)

	mux.Handle("/console", gql.GraphqlwsHandler)

	handler := cors.Default().Handler(mux)

	log.Println(fmt.Sprintf("Serve at %s...", config.GetPort()))
	http.ListenAndServe(config.GetPort(), handler)
}

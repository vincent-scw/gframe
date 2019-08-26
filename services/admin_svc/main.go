package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	r "github.com/vincent-scw/gframe/redisctl"

	"github.com/vincent-scw/gframe/admin_svc/config"
	gql "github.com/vincent-scw/gframe/admin_svc/graphql"
	"github.com/vincent-scw/gframe/admin_svc/subscriber"
)

func main() {
	log.Println("Starting admin service...")

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I am good."))
	})

	h := handler.New(&handler.Config{
		Schema: &gql.Schema,
		Pretty: true,
	})
	mux.Handle("/graphql", h)

	mux.Handle("/metrics", promhttp.Handler())

	mux.Handle("/console", gql.GraphqlwsHandler)

	handler := cors.Default().Handler(mux)

	pubsub := r.NewRedisClient(config.GetRedisServer())
	defer pubsub.Close()

	log.Println("Subscribe to Redis...")
	go subscriber.SubscribePlayer(pubsub, gql.Broadcast)
	go subscriber.SubscribeGroup(pubsub, gql.Broadcast)

	log.Println(fmt.Sprintf("Serve at %s...", config.GetPort()))
	http.ListenAndServe(config.GetPort(), handler)
}

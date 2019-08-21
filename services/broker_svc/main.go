package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	e "github.com/vincent-scw/gframe/contracts"
)

func main() {
	log.Println("Starting broker service...")

	// Set default configurations
	viper.SetDefault("WEB_PORT", 8443)
	viper.SetDefault("RPC_PORT", 8543)
	viper.SetDefault("REDIS_SERVER", "localhost:6379")
	viper.SetDefault("KAFKA_BROKERS", []string{"localhost:9092"})

	viper.AutomaticEnv() // automatically bind env

	webPort := viper.GetInt("WEB_PORT")
	rpcPort := viper.GetInt("RPC_PORT")

	go serveWeb(webPort)

	ctx, cancel := context.WithCancel(context.Background())
	serveRPC(rpcPort)
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

func serveRPC(port int) {
	handler := newReceptionHandler()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	e.RegisterUserReceptionServer(s, handler)
	log.Printf("Listen to RPC at port %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveWeb(port int) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am healthy.")
	})
	log.Println(fmt.Sprintf("Serve start at %d.", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

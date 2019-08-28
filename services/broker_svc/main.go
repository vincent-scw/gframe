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

	"google.golang.org/grpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	e "github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/broker_svc/config"
)

func main() {
	log.Println("Starting broker service...")

	go serveWeb(config.GetWebPort())

	ctx, cancel := context.WithCancel(context.Background())
	serveRPC(config.GetRPCPort())
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

func serveRPC(port string) {
	handler := newReceptionHandler()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	e.RegisterUserReceptionServer(s, handler)
	log.Printf("Listen to RPC at port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveWeb(port string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am healthy.")
	})
	log.Println(fmt.Sprintf("Serve start at %s.", port))
	log.Fatal(http.ListenAndServe(port, nil))
}

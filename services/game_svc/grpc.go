package main

import (
	"log"

	"google.golang.org/grpc"

	"github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/game_svc/config"
)

type brokerRPCWrapper struct {
	client     contracts.UserReceptionClient
	clientConn *grpc.ClientConn
}

func newBrokerRPCWrapper() *brokerRPCWrapper {
	conn, err := grpc.Dial(config.GetBrokerRPC(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	rpc := contracts.NewUserReceptionClient(conn)

	wrapper := brokerRPCWrapper{client: rpc, clientConn: conn}
	return &wrapper
}

func (wrapper *brokerRPCWrapper) Close() {
	wrapper.clientConn.Close()
}

package main

import (
	"context"
	"encoding/json"

	g "github.com/vincent-scw/gframe/broker_svc/game"
	"github.com/vincent-scw/gframe/broker_svc/singleton"
	e "github.com/vincent-scw/gframe/contracts"
)

type receptionHandler struct {
	matching *g.Matching
}

func newReceptionHandler() *receptionHandler {
	handler := &receptionHandler{}
	// start matching
	handler.matching = g.NewMatching(2, 1000, 30)
	handler.matching.Formed = func(g *g.Group) {
		value, _ := json.Marshal(g)
		go singleton.RedisPublish(e.GroupChannel, string(value))
	}
	return handler
}

func (handler *receptionHandler) Checkin(ctx context.Context, user *e.User) (*e.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.RedisPublish(e.PlayerChannel, string(b))

	result := handler.matching.AddToGroup(*user)
	return &e.ReceptionResponse{Acknowledged: result}, nil
}

func (handler *receptionHandler) Checkout(ctx context.Context, user *e.User) (*e.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.RedisPublish(e.PlayerChannel, string(b))

	return &e.ReceptionResponse{Acknowledged: true}, nil
}

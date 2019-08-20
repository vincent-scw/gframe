package main

import (
	"context"
	"encoding/json"

	g "github.com/vincent-scw/gframe/broker_svc/game"
	"github.com/vincent-scw/gframe/broker_svc/singleton"
	e "github.com/vincent-scw/gframe/events"
)

type receptionHandler struct {
	matching *g.Matching
}

func newReceptionHandler() *receptionHandler {
	handler := &receptionHandler{}
	// start matching
	handler.matching = g.NewMatching(2, 1000, 30)

	return handler
}

func (handler *receptionHandler) CheckIn(ctx context.Context, user *e.User) (*e.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.GetPubSubClient().Publish(e.PlayerChannel, string(b))

	result := handler.matching.AddToGroup(*user)
	return &e.ReceptionResponse{Acknowledged: result}, nil
}

func (handler *receptionHandler) CheckOut(ctx context.Context, user *e.User) (*e.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.GetPubSubClient().Publish(e.PlayerChannel, string(b))

	result := handler.matching.AddToGroup(*user)
	return &e.ReceptionResponse{Acknowledged: result}, nil
}

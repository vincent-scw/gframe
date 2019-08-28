package main

import (
	"context"
	"encoding/json"
	"fmt"

	g "github.com/vincent-scw/gframe/broker_svc/game"
	"github.com/vincent-scw/gframe/broker_svc/singleton"
	c "github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/redisctl"
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

		go onFormedGroup(value, g)
	}
	return handler
}

func (handler *receptionHandler) Checkin(ctx context.Context, user *c.User) (*c.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.GetRedisClient().Publish(c.PlayerChannel, string(b))

	result := handler.matching.AddToGroup(*user)
	return &c.ReceptionResponse{Acknowledged: result}, nil
}

func (handler *receptionHandler) Checkout(ctx context.Context, user *c.User) (*c.ReceptionResponse, error) {
	// Send to Redis pub/sub
	b, _ := json.Marshal(user)
	go singleton.GetRedisClient().Publish(c.PlayerChannel, string(b))

	return &c.ReceptionResponse{Acknowledged: true}, nil
}

func onFormedGroup(value []byte, g *g.Group) {
	redisCli := singleton.GetRedisClient()
	redisCli.SetCache(fmt.Sprintf(redisctl.GROUP_FORMAT, "1", g.ID), string(value), 0)
	redisCli.Publish(c.GroupChannel, string(value))
}

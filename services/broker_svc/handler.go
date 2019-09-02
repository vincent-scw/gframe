package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"

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

func (handler *receptionHandler) Handle(message *sarama.ConsumerMessage) bool {
	// Send to Redis pub/sub
	go singleton.GetRedisClient().Publish(c.PlayerChannel, string(message.Value))

	event := &c.UserEvent{}
	err := json.Unmarshal(message.Value, event)
	if err != nil {
		log.Println("Unable to unmarshal to UserEvent from Kafka message.")
		return false
	}

	switch event.Status {
	case c.UserEvent_In:
		return handler.matching.AddToGroup(*event.User)
	case c.UserEvent_Out:
		break
	default:
		log.Println("Not supported Kafka message.")
		return false
	}
	return true
}

func onFormedGroup(value []byte, g *g.Group) {
	redisCli := singleton.GetRedisClient()
	redisCli.SetCache(fmt.Sprintf(redisctl.GroupFormat, "1", g.ID), string(value), 30*60*1000)
	redisCli.Publish(c.GroupChannel, string(value))
}

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
	go singleton.GetRedisClient().Publish(redisctl.PlayerChannel, string(message.Value))

	event := &c.UserEvent{}
	err := json.Unmarshal(message.Value, event)
	if err != nil {
		log.Println("Unable to unmarshal to UserEvent from Kafka message.")
		return false
	}

	switch event.Type {
	case c.EventType_In:
		return handler.matching.AddToGroup(*event.User)
	case c.EventType_Out:
		break
	default:
		log.Println("Not supported Kafka message.")
		return false
	}
	return true
}

func onFormedGroup(value []byte, g *g.Group) {
	redisCli := singleton.GetRedisClient()
	// store to cache
	redisCli.SetCache(fmt.Sprintf(redisctl.GroupFormat, "1", g.Id), string(value), 30*60*1000)
	// publish to clients
	redisCli.Publish(redisctl.GroupChannel, string(value))
	// start to wait for player interaction
}

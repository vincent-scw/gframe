package main

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	g "github.com/vincent-scw/gframe/broker_svc/game"
	e "github.com/vincent-scw/gframe/events"
	"github.com/vincent-scw/gframe/broker_svc/singleton"
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

func (handler *receptionHandler) Handle(message *sarama.ConsumerMessage) bool {
	// Send to Redis pub/sub
	go singleton.GetPubSubClient().Publish(e.PlayerChannel, string(message.Value))

	event := &e.UserEvent{}
	err := json.Unmarshal(message.Value, event)
	if err != nil {
		log.Println("Unable to unmarshal to UserEvent from Kafka message.")
		return false
	}

	switch event.Type {
	case e.UserEventIn:
		return handler.matching.AddToGroup(event.User)
	case e.UserEventOut:
		break
	default:
		log.Println("Not supported Kafka message.")
		return false
	}
	return true
}

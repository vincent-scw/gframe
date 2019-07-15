package main

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	e "github.com/vincent-scw/gframe/kafka/events"
)

type receptionHandler struct {
	event chan e.UserEvent
}

func newReceptionHandler() *receptionHandler {
	return &receptionHandler{}
}

func (handler *receptionHandler) Handle(message *sarama.ConsumerMessage) bool {
	event := &e.UserEvent{}
	err := json.Unmarshal(message.Value, event)
	if err != nil {
		log.Println("Unable to unmarshal to UserEvent from Kafka message.")
		return false
	}

	switch event.Type {
	case e.EventIn:
		break
	case e.EventOut:
		break
	default:
		log.Println("Not supported Kafka message.")
		return false
	}
	return true
}

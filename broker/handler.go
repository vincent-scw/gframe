package main

import (
	"github.com/Shopify/sarama"
	_ "github.com/vincent-scw/gframe/kafka/events"
)

type playerReceptionHandler struct {
}

func (handler *playerReceptionHandler) Handle(*sarama.ConsumerMessage) bool {
	return true
}

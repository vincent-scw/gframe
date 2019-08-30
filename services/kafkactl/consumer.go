package kafkactl

import (
	"log"

	"github.com/Shopify/sarama"
)

// Consumer represents a Sarama consumer group
type Consumer struct {
	Ready   chan bool
	Handler MessageHandler
}

// MessageHandler handles kafka message
type MessageHandler interface {
	Handle(*sarama.ConsumerMessage) bool
}

// Setup is run at the begining of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Kafka message claimed: value = %s, topic = %s", string(message.Value), message.Topic)
		if consumer.Handler != nil && consumer.Handler.Handle(message) {
			session.MarkMessage(message, "")
		}
	}

	return nil
}

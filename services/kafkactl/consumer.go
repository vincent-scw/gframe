package kafkactl

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

// Consumer represents a consumer
type Consumer struct {
	Ready   chan bool
	Handler MessageHandler
	Context context.Context
}

// MessageHandler handles kafka message
type MessageHandler interface {
	Handle(*sarama.ConsumerMessage) bool
}

// ConsumerGroup represents a Sarama  consumer group
type ConsumerGroup struct {
	client sarama.ConsumerGroup
}

// NewConsumerGroup creates kafka consumer
func NewConsumerGroup(brokers []string, topic KafkaTopic, consumer *Consumer) (*ConsumerGroup, error) {
	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, fmt.Sprintf("%s_broker", topic.name), config)

	if err != nil {
		return nil, err
	}

	go client.Consume(consumer.Context, []string{topic.name}, consumer)
	return &ConsumerGroup{client: client}, nil
}

// Close release
func (c *ConsumerGroup) Close() {
	if c.client != nil {
		c.client.Close()
	}
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

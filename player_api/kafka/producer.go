package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// Producer is the kafka message producer
type Producer struct {
}

// NewProducer returns a new producer
func NewProducer() *Producer {
	p := &Producer{}
	return p
}

// Emit sends a message to kafka
func (p *Producer) Emit(key string, msg string) {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"40.83.99.7:9092"},
		Topic:    "player",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()
	writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(msg),
		})

	fmt.Println("message emitted")
}

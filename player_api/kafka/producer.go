package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// Producer is the kafka message producer
type Producer struct {
}

// TextMessage is a pure text message
type TextMessage struct {
	Key     string `json:"key"`
	Content string `json:"content"`
}

// NewProducer returns a new producer
func NewProducer() *Producer {
	p := &Producer{}
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"40.83.99.7:9092"},
		Topic:    "player",
		Balancer: &kafka.LeastBytes{},
	})
	return p
}

// Emit sends a message to kafka
func (p *Producer) Emit(msg *TextMessage) {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(msg.Key),
			Value: []byte(msg.Content),
		})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("message emitted")
}

// EmitMulti emits multiple messages
func (p *Producer) EmitMulti(msgs []TextMessage) {
	messages := make([]kafka.Message, len(msgs))
	for i, msg := range msgs {
		messages[i] = kafka.Message{
			Key:   []byte(msg.Key),
			Value: []byte(msg.Content),
		}
	}

	err := writer.WriteMessages(context.Background(), messages...)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("messages emitted")
}

// Dispose releases resources
func (p *Producer) Dispose() {
	writer.Close()
}

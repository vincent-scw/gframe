package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// Producer is the kafka message producer
type Producer struct {
}

// KeyDef is an interface
type KeyDef interface {
	DefKey() string
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
func (p *Producer) Emit(model KeyDef) (err error) {
	key := model.DefKey()
	value, err := json.Marshal(model)
	if err != nil {
		return
	}

	msg := kafka.Message{Key: []byte(key), Value: value}
	err = p.emit(msg)
	return
}

func (p *Producer) emit(msg kafka.Message) error {
	err := writer.WriteMessages(context.Background(), msg)

	if err != nil {
		log.Fatalln(err)
		return err
	}
	fmt.Println("message emitted")
	return err
}

// Dispose releases resources
func (p *Producer) Dispose() {
	writer.Close()
}

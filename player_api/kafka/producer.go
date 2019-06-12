package kafka

import (
	"fmt"
	"log"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	brokers             = []string{"40.83.99.7:9092"}
	topic   goka.Stream = "player"
	group   goka.Group  = "test-group"
)

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
	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()
	err = emitter.EmitSync(key, msg)
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
	fmt.Println("message emitted")
}

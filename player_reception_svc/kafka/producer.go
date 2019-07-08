package kafka

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

// Producer is the kafka message producer
type Producer struct {
	PlayerProducer sarama.SyncProducer
}

// KeyDef is an interface
type KeyDef interface {
	DefKey() string
}

// NewProducer returns a new producer
func NewProducer() *Producer {
	p := &Producer{
		PlayerProducer: newPlayerProducer([]string{"40.83.112.48:9092"}),
	}

	return p
}

// Emit sends a message to kafka
func (p *Producer) Emit(model KeyDef) (err error) {
	key := model.DefKey()
	value, err := json.Marshal(model)
	if err != nil {
		return
	}

	_, _, err = p.PlayerProducer.SendMessage(&sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: "player",
		Value: sarama.ByteEncoder(value),
	})

	return
}

func newPlayerProducer(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

// Dispose releases resources
func (p *Producer) Dispose() error {
	if err := p.PlayerProducer.Close(); err != nil {
		log.Println("Failed to shut down player producer cleanly", err)
	}
	return nil
}

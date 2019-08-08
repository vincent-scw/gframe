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

// KeyDef is key definition
type KeyDef interface {
	DefKey() string
}

// NewProducer returns a new producer
func NewProducer(brokers []string) (*Producer, error) {
	pp, err := newPlayerProducer(brokers)
	p := &Producer{
		PlayerProducer: pp,
	}

	return p, err
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

func newPlayerProducer(brokerList []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokerList, config)
}

// Dispose releases resources
func (p *Producer) Dispose() error {
	if err := p.PlayerProducer.Close(); err != nil {
		log.Println("Failed to shut down player producer cleanly", err)
	}
	return nil
}

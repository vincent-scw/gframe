package kafkactl

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

// Producer is the kafka message producer
type Producer struct {
	syncProducer  sarama.SyncProducer
	topicSettings KafkaTopic
}

// KeyDef is key definition
type KeyDef interface {
	GetID() string
}

// NewProducer returns a new producer
func NewProducer(brokers []string, topic KafkaTopic) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	if topic.strategy == hash {
		config.Producer.Partitioner = sarama.NewHashPartitioner
	}

	kp, err := sarama.NewSyncProducer(brokers, config)
	p := &Producer{
		syncProducer:  kp,
		topicSettings: topic,
	}

	return p, err
}

// Emit sends a message to kafka
func (p *Producer) Emit(model KeyDef) (err error) {
	key := model.GetID()
	value, err := json.Marshal(model)
	if err != nil {
		return
	}

	msg := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: p.topicSettings.name,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = p.syncProducer.SendMessage(msg)

	if err == nil {
		log.Printf("send to kafka succeed: %s", string(value))
	} else {
		log.Fatalf("error talk to kafka: %v", err)
	}
	return
}

// Dispose releases resources
func (p *Producer) Dispose() {
	p.syncProducer.Close()
}

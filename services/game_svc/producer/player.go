package producer

import (
	"log"
	"time"

	k "github.com/vincent-scw/gframe/kafkactl"
	"github.com/vincent-scw/gframe/util"

	"github.com/vincent-scw/gframe/game_svc/config"
)

// PlayerEventProducer producer
type PlayerEventProducer struct {
	kafka *k.Producer
}

// NewPlayerEventProducer generates a new producer
func NewPlayerEventProducer() *PlayerEventProducer {
	p := PlayerEventProducer{}
	err := util.WithRetry(10, 2*time.Second, func() (err error) {
		p.kafka, err = k.NewProducer(config.GetKafkaBrokers(), k.TopicPlayer)
		return
	})
	
	if err != nil {
		log.Panic(err)
	}
	return &p
}

// Emit player to kafka
func (p *PlayerEventProducer) Emit(msg k.KeyDef) error {
	err := p.kafka.Emit(msg); 
	if err != nil {
		log.Fatalf("emitting player to kafka error: %v", err)
	}
	return err
}

// Close dispose things
func (p *PlayerEventProducer) Close() {
	if p.kafka != nil {
		p.kafka.Dispose()
	}
}


package producer

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

func runEmitter() {
	emitter, err := goka.NewEmitter(brokers, topc, new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()
	err = emitter.EmitSync("some-key", "some-value")
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
	fmt.Println("message emitted")
}

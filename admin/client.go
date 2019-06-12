package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"40.83.99.7:9092"},
		Topic:   "player",
		GroupID: "admin-group",
	})

	defer r.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at partition/offset %v/%v: %s = %s\n",
			m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

}

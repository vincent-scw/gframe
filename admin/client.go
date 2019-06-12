package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"40.83.99.7:9092"},
		Topic:   "player",
		GroupID: "admin-group",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error %v", err)
			break
		}
		fmt.Printf("message at partition/offset %v/%v: %s = %s\n",
			m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	
}

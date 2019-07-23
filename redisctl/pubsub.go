package redisctl

import (
	"fmt"

	"github.com/go-redis/redis"
)

// PubSubClient represents a Redis pub-sub client
type PubSubClient struct {
	redisdb *redis.Client
}

// Handle is a function to handle received content
type Handle func(string)

// NewPubSubClient creates a pub-sub client
func NewPubSubClient(addr ...string) *PubSubClient {
	cli := &PubSubClient{}
	cli.redisdb = redis.NewClient(&redis.Options{
		Addr:     addr[0],
		Password: "",
		DB:       0,
	})

	return cli
}

// Publish publishes content to channel
func (cli *PubSubClient) Publish(channel string, content string) {
	err := cli.redisdb.Publish(channel, content).Err()
	if err != nil {
		panic(err)
	}
}

// Subscribe subscribes a channel
func (cli *PubSubClient) Subscribe(channel string, handle Handle) {
	pubsub := cli.redisdb.Subscribe(channel)
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		if handle != nil {
			handle(msg.Payload)
		}
	}
}

// Close releases resources
func (cli *PubSubClient) Close() {
	if err := cli.redisdb.Close(); err != nil {
		panic(err)
	}
}

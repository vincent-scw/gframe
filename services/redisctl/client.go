package redisctl

import (
	"log"

	"github.com/go-redis/redis"
)

// RedisClient represents a Redis client
type RedisClient struct {
	redisdb *redis.Client
}

// Handle is a function to handle received content
type Handle func(string) string

// NewRedisClient creates a redis client
func NewRedisClient(addr ...string) *RedisClient {
	cli := &RedisClient{}
	cli.redisdb = redis.NewClient(&redis.Options{
		Addr:     addr[0],
		Password: "",
		DB:       0,
	})

	return cli
}

// Publish publishes content to channel
func (cli *RedisClient) Publish(channel string, content string) {
	log.Printf("Publish event to Redis channel [%s] %s", channel, content)
	err := cli.redisdb.Publish(channel, content).Err()
	if err != nil {
		panic(err)
	}
}

// Subscribe subscribes a channel
func (cli *RedisClient) Subscribe(channel string, handles ...Handle) {
	pubsub := cli.redisdb.Subscribe(channel)
	ch := pubsub.Channel()

	for msg := range ch {
		log.Println(msg.Channel, msg.Payload)
		if handles != nil {
			tmp := msg.Payload
			for _, handle := range handles {
				tmp = handle(tmp)
			}
		}
	}
}

// Close releases resources
func (cli *RedisClient) Close() {
	if err := cli.redisdb.Close(); err != nil {
		panic(err)
	}
}

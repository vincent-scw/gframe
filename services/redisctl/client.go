package redisctl

import (
	"log"
	"time"

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

// SetCache set cache to Redis
func (cli *RedisClient) SetCache(key string, value interface{}, exp time.Duration) {
	err := cli.redisdb.Set(key, value, exp).Err()
	if err != nil {
		log.Printf("Set to Redis error %v", err)
	}
}

// GetCache get cache from Redis
func (cli *RedisClient) GetCache(key string) (string, error) {
	v, err := cli.redisdb.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		log.Printf("Read from Redis error %v", err)
	}
	return v, err
}

// Increment a integer value
func (cli *RedisClient) Increment(key string) int64 {
	result, err := cli.redisdb.Incr(key).Result()
	if err != nil {
		log.Printf("Increment error %v", err)
	}
	return result
}

// Decrement a integer value
func (cli *RedisClient) Decrement(key string) int64 {
	result, err := cli.redisdb.Decr(key).Result()
	if err != nil {
		log.Printf("Decrement error %v", err)
	}
	return result
}

// PushToList use RPush to a redis list
func (cli *RedisClient) PushToList(key string, value string) {
	err := cli.redisdb.RPush(key, value).Err()
	if err != nil {
		log.Printf("Push to list error %v", err)
	}
}

// GetAllFromList pops from redis list
func (cli *RedisClient) GetAllFromList(key string) []string {
	result, err := cli.redisdb.LRange(key, 0, cli.GetListLength(key)).Result()
	if err != nil {
		log.Printf("Pop from list error %v", err)
	}
	return result
}

// GetListLength returns the length of redis list
func (cli *RedisClient) GetListLength(key string) int64 {
	length, err := cli.redisdb.LLen(key).Result()
	if err != nil {
		log.Printf("Get length from list error %v", err)
	}
	return length
}

// Close releases resources
func (cli *RedisClient) Close() {
	if err := cli.redisdb.Close(); err != nil {
		panic(err)
	}
}

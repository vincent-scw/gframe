package singleton

import (
	"sync"

	"github.com/spf13/viper"

	r "github.com/vincent-scw/gframe/redisctl"
)

var (
	once        sync.Once
	redisClient *r.RedisClient
)

// GetRedisClient returns redisClient
func GetRedisClient() *r.RedisClient {
	if viper.GetString("REDIS_SERVER") == "" {
		return nil
	}
	once.Do(func() {
		redisClient = r.NewRedisClient(viper.GetString("REDIS_SERVER"))
	})

	return redisClient
}

// RedisPublish message to Redis
func RedisPublish(channel string, content string) {
	pubSubClient := GetRedisClient()
	if pubSubClient != nil {
		pubSubClient.Publish(channel, content)
	}
}

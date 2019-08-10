package singleton

import (
	"sync"

	"github.com/spf13/viper"

	r "github.com/vincent-scw/gframe/redisctl"
)

var (
	once         sync.Once
	pubsubClient *r.PubSubClient
)

// GetPubSubClient returns pubsubclient
func GetPubSubClient() *r.PubSubClient {
	if viper.GetString("REDIS_SERVER") == "" {
		return nil
	}
	once.Do(func() {
		pubsubClient = r.NewPubSubClient(viper.GetString("REDIS_SERVER"))
	})

	return pubsubClient
}

// RedisPublish message to Redis
func RedisPublish(channel string, content string) {
	pubSubClient := GetPubSubClient()
	if pubSubClient != nil {
		pubSubClient.Publish(channel, content)
	}
}

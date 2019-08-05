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
	once.Do(func() {
		pubsubClient = r.NewPubSubClient(viper.GetString("REDIS_SERVER"))
	})

	return pubsubClient
}

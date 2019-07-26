package singleton

import (
	"sync"

	r "github.com/vincent-scw/gframe/redisctl"
)

var (
	once sync.Once
	pubsubClient *r.PubSubClient
)

// GetPubSubClient returns pubsubclient
func GetPubSubClient() *r.PubSubClient {
	once.Do(func() {
		pubsubClient = r.NewPubSubClient("40.83.112.48:6379")
	})

	return pubsubClient
}
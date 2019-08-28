package singleton

import (
	"sync"

	r "github.com/vincent-scw/gframe/redisctl"
	"github.com/vincent-scw/gframe/broker_svc/config"
)

var (
	once        sync.Once
	redisClient *r.RedisClient
)

// GetRedisClient returns redisClient
func GetRedisClient() *r.RedisClient {
	if config.GetRedisServer() == "" {
		return nil
	}
	once.Do(func() {
		redisClient = r.NewRedisClient(config.GetRedisServer())
	})

	return redisClient
}

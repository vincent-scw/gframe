package subscriber

import (
	r "github.com/vincent-scw/gframe/redisctl"
	"github.com/vincent-scw/gframe/game_svc/singleton"
)

// ChannelSubscriber interface
type ChannelSubscriber interface {
	subscribe(*r.RedisClient)
}

// ChannelSubscribers listens channels
type ChannelSubscribers struct {
	client *r.RedisClient
	subscribers []ChannelSubscriber
}

// NewChannelSubscribers listens from Redis channels
func NewChannelSubscribers(subs ...ChannelSubscriber) *ChannelSubscribers {
	sub := ChannelSubscribers{ 
		client: singleton.GetRedisClient(),
		subscribers: subs,
	}
	
	return &sub
}

// StartSubscribing start subscribing
func (subs *ChannelSubscribers) StartSubscribing() {
	for _, sub := range subs.subscribers {
		go sub.subscribe(subs.client)
	}
}

package game

import (
	c "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"
)

// Service is service
type Service struct {
	redisClient *r.RedisClient
}

// NewService returns game service
func NewService(redis *r.RedisClient) *Service {
	svc := Service{redisClient: redis}
	return &svc
}

// Play a game
func (svc *Service) Play(game *c.Game) {

}

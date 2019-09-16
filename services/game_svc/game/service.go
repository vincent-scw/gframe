package game

import (
	"fmt"
	"log"

	c "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"
	u "github.com/vincent-scw/gframe/util"

	"github.com/vincent-scw/gframe/game_svc/singleton"
)

// Service is service
type Service struct {
	redisClient *r.RedisClient
}

// NewService returns game service
func NewService() *Service {
	svc := Service{redisClient: singleton.GetRedisClient()}
	return &svc
}

// Play a game
func (svc *Service) Play(game *c.GameEvent) {
	if game.Group == nil {
		log.Printf("data with error: no group in %s", string(u.ToJSON(game)))
		return
	}
	gameKey := fmt.Sprintf(r.GameEventFormat, game.Group.Id)
	countKey := fmt.Sprintf(r.GameEventCountFormat, game.Group.Id)
	svc.redisClient.PushToList(gameKey, string(u.ToJSON(game.Move)))
	count := svc.redisClient.Increment(countKey)
	// TODO: get 2 move, shall be configurable
	if count == 2 {
		list := svc.redisClient.GetAllFromList(gameKey)
		moves := toMovesList(list)
		result := judge(moves)
		svc.redisClient.Publish(r.GameChannel, string(u.ToJSON(result)))
	}
}

func toMovesList(strs []string) []*c.Move {
	var ret []*c.Move
	for _, g := range strs {
		p := c.Move{}
		u.ToModel([]byte(g), &p)
		ret = append(ret, &p)
	}
	return ret
}

func judge(moves []*c.Move) *c.Result {
	if len(moves) != 2 {
		log.Print("wrong round.")
		return nil
	}
	result := c.Result{Moves: moves}
	p1 := moves[0]
	p2 := moves[1]
	diff := p1.Shape - p2.Shape
	if diff == 1 || diff == -2 {
		// p1 win
		result.Winner = 0
	} else if diff == 0 {
		// draw game
		result.Winner = -1
	} else {
		// p2 win
		result.Winner = 1
	}
	return &result
}

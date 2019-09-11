package game

import (
	"encoding/json"
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
	gameKey := fmt.Sprintf(r.GameEventFormat, game.Group.Id)
	countKey := fmt.Sprintf(r.GameEventCountFormat, game.Group.Id)
	svc.redisClient.PushToList(gameKey, string(u.ToJSON(game.Play)))
	count := svc.redisClient.Increment(countKey)
	// TODO: get 2 play, shall be configurable
	if count == 2 {
		list := svc.redisClient.GetAllFromList(gameKey)
		plays := toPlayingList(list)
		result := judge(plays)
		resultStr, _ := json.Marshal(result)
		svc.redisClient.Publish(r.GameChannel, string(resultStr))
	}
}

func toPlayingList(strs []string) []*c.Playing {
	var ret []*c.Playing
	for _, g := range strs {
		p := c.Playing{}
		json.Unmarshal([]byte(g), &p)
		ret = append(ret, &p)
	}
	return ret
}

func judge(plays []*c.Playing) *c.Result {
	if len(plays) != 2 {
		log.Print("wrong round.")
		return nil
	}
	result := c.Result{Plays: plays}
	p1 := plays[0]
	p2 := plays[1]
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

package subscriber

import (
	"encoding/json"
	"log"

	c "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"

	conn "github.com/vincent-scw/gframe/game_svc/connection"
)

// PlaySubscriber subscribes group channel
type PlaySubscriber struct {
	hub *conn.Hub
}

// NewPlaySubscriber returns PlaySubscriber
func NewPlaySubscriber(hub *conn.Hub) *PlaySubscriber {
	return &PlaySubscriber{hub: hub}
}

func (sub *PlaySubscriber) subscribe(client *r.RedisClient) {
	client.Subscribe(r.GameChannel, sub.handlePlay)
}

func (sub *PlaySubscriber) handlePlay(msg string) string {
	event := &c.Result{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range event.Plays {
		// try send to player client
		sub.hub.SendToClient(p.Player.Id, conn.NewMessage(conn.Game, []byte(msg)))
	}

	return msg
}
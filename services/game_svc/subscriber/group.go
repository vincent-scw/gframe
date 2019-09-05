package subscriber

import (
	"encoding/json"
	"log"

	c "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"

	conn "github.com/vincent-scw/gframe/game_svc/connection"
)

// GroupSubscriber subscribes group channel
type GroupSubscriber struct {
	hub *conn.Hub
}

// NewGroupSubscriber returns GroupSubscriber
func NewGroupSubscriber(hub *conn.Hub) *GroupSubscriber {
	return &GroupSubscriber{hub: hub}
}

func (sub *GroupSubscriber) subscribe(client *r.RedisClient) {
	client.Subscribe(c.GroupChannel, sub.handleGroup)
}

func (sub *GroupSubscriber) handleGroup(msg string) string {
	event := &c.GroupInfo{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	switch event.Status {
	case c.GroupFormed:
		for _, p := range event.Players {
			// try send to player client
			sub.hub.SendToClient(p.ID, conn.NewMessage(conn.Group, []byte(msg)))
		}
	default:
		break
	}

	return msg
}

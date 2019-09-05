package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	c "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"
)

// SubscribePlayer subscribes player channel
func SubscribePlayer(client *r.RedisClient, foo func(formattedMsg string) string) {
	client.Subscribe(r.PlayerChannel, handlePlayer, foo)
}

func handlePlayer(msg string) string {
	event := &c.UserEvent{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	var formatted string
	switch event.Type {
	case c.EventType_In:
		formatted = withTime(fmt.Sprintf("Player %s joined the game.", withColor(event.User.Name, yellow)))
	case c.EventType_Out:
		formatted = withTime(fmt.Sprintf("Player %s left the game.", withColor(event.User.Name, yellow)))
	default:
		formatted = msg
	}

	return formatted
}

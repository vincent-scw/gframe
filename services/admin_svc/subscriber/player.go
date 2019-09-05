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
	client.Subscribe(r.PlayerChannel, func(msg string) string {
		event := &c.UserEvent{}
		err := json.Unmarshal([]byte(msg), event)
		if err != nil {
			log.Fatal(err)
		}

		var formatted string
		switch event.Type {
		case c.EventType_In:
			formatted = withTime(
				fmt.Sprintf("Player %s joined the game. (Total: %d)",
					withColor(event.User.Name, yellow),
					client.Increment(r.PlayerCountFormat)))
		case c.EventType_Out:
			formatted = withTime(
				fmt.Sprintf("Player %s left the game. (Total: %d)",
					withColor(event.User.Name, yellow),
					client.Decrement(r.PlayerCountFormat)))
		default:
			formatted = msg
		}

		return formatted
	}, foo)
}

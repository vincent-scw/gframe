package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

// SubscribePlayer subscribes player channel
func SubscribePlayer(client *r.PubSubClient, foo func(formattedMsg string) string) {
	client.Subscribe(e.PlayerChannel, handlePlayer, foo)
}

func handlePlayer(msg string) string {
	event := &e.UserEvent{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	var formatted string
	switch event.Type {
	case e.UserEventIn:
		formatted = fmt.Sprintf("<i>INFO</i> Player %s joined the game.", 
			withColor(event.Name, green))
	case e.UserEventOut:
		formatted = fmt.Sprintf("<i>INFO</i> Player %s left the game.", 
			withColor(event.Name, red))
	default:
		formatted = msg
	}

	return formatted
}
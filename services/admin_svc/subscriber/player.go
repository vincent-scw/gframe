package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	e "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"
)

// SubscribePlayer subscribes player channel
func SubscribePlayer(client *r.PubSubClient, foo func(formattedMsg string) string) {
	client.Subscribe(e.PlayerChannel, handlePlayer, foo)
}

func handlePlayer(msg string) string {
	event := &e.User{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	var formatted string
	switch event.Status {
	case e.User_In:
		formatted = withTime(fmt.Sprintf("Player %s joined the game.", withColor(event.Name, yellow)))
	case e.User_Out:
		formatted = withTime(fmt.Sprintf("Player %s left the game.", withColor(event.Name, yellow)))
	default:
		formatted = msg
	}

	return formatted
}

package subscriber

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

// SubscribeGroup subscribes group channel
func SubscribeGroup(client *r.PubSubClient, foo func(formattedMsg string) string) {
	client.Subscribe(e.GroupChannel, handleGroup, foo)
}

func handleGroup(msg string) string {
	event := &e.GroupInfo{}
	err := json.Unmarshal([]byte(msg), event)
	if err != nil {
		log.Fatal(err)
	}

	var formatted string
	switch event.Status {
	case e.GroupFormed:
		formatted = withTime(fmt.Sprintf(
			"Group %s has been formed with players %s",
			withColor(event.ID, yellow), withColor(playersToString(event.Players), yellow)))
	default:
		formatted = msg
	}

	return formatted
}

func playersToString(players []e.User) string {
	valuesText := []string{}
	for _, p := range players {
		valuesText = append(valuesText, p.Name)
	}
	return strings.Join(valuesText, ", ")
}

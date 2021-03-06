package subscriber

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	e "github.com/vincent-scw/gframe/contracts"
	r "github.com/vincent-scw/gframe/redisctl"
)

// SubscribeGroup subscribes group channel
func SubscribeGroup(client *r.RedisClient, foo func(formattedMsg string) string) {
	client.Subscribe(r.GroupChannel, func(msg string) string {
		event := &e.GroupInfo{}
		err := json.Unmarshal([]byte(msg), event)
		if err != nil {
			log.Fatal(err)
		}

		var formatted string
		switch event.Status {
		case e.GroupStatus_Formed:
			formatted = withTime(fmt.Sprintf(
				"Group %s (total: %d) has been formed with players %s.",
				withColor(event.Id, yellow),
				client.Increment(r.GroupCountFormat),
				withColor(playersToString(event.Players), yellow)))
		default:
			formatted = msg
		}

		return formatted
	}, foo)
}

func playersToString(players []*e.User) string {
	valuesText := []string{}
	for _, p := range players {
		valuesText = append(valuesText, p.Name)
	}
	return strings.Join(valuesText, ", ")
}

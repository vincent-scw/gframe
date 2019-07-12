package game

import (
	"testing"

	"github.com/vincent-scw/gframe/kafka/events"
)

func TestMatching(t *testing.T) {
	matching := NewMatching()
	player := events.User{ID: "ABC", Name: "ABC"}
	matching.AddToGroup(&player)

	if len(matching.Groups) != 1 {
		t.Errorf("One group should be created.")
	}
}

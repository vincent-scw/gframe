package game

import (
	"fmt"
	"testing"

	"github.com/vincent-scw/gframe/kafka/events"
)

func TestMatching(t *testing.T) {
	matching := NewMatching(2, 10, 1)

	var player events.User
	count := 0
	for i := 0; i < 20; i++ {
		id := fmt.Sprintf("User%d", i)
		player = events.User{ID: id, Name: id}
		if matching.AddToGroup(player) {
			count++
		}
	}

	for k, g := range matching.Groups {
		t.Logf("Group %s has players %s, %s", k, g.Players[0].Name, g.Players[1].Name)
	}

	if len(matching.Groups) != 10 {
		t.Errorf("10 groups should be created, but was %d.", len(matching.Groups))
	}

	if count != 20 {
		t.Errorf("All users should be added to group, but only %d added.", count)
	}
}

func BenchmarkMatching(b *testing.B) {
	matching := NewMatching(2, b.N/2+1, 1)

	var player events.User
	for i := 0; i < b.N; i++ {
		id := fmt.Sprintf("User%d", i)
		player = events.User{ID: id, Name: id}
		matching.AddToGroup(player)
	}
	b.Logf("%d groups have been formed.", len(matching.Groups))
}

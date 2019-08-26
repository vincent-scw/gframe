package game

import (
	"fmt"
	"testing"
	"time"

	"github.com/vincent-scw/gframe/contracts"
)

func TestMatching(t *testing.T) {
	var totalUsers = 1000
	var groupSize = 2
	var totalGroups = totalUsers / groupSize

	matching := NewMatching(groupSize, totalGroups, 1)
	var player contracts.User
	count := 0
	for i := 0; i < totalUsers; i++ {
		id := fmt.Sprintf("User%d", i)
		player = contracts.User{Id: id, Name: id}
		if matching.AddToGroup(player) {
			count++
		}
	}

	time.Sleep(time.Millisecond * time.Duration(100))

	for k, g := range matching.Groups {
		if len(g.Players) != groupSize {
			t.Errorf("Group %s has %d players", k, len(g.Players))
		}
	}

	if len(matching.Groups) != totalGroups {
		t.Errorf("%d groups should be created, but was %d.", totalGroups, len(matching.Groups))
	}

	if count != totalUsers {
		t.Errorf("All users should be added to group, but only %d added.", count)
	}
}

func BenchmarkMatching(b *testing.B) {
	matching := NewMatching(2, b.N/2+1, 1)
	b.Logf("%d players to be injected.", b.N)
	var player contracts.User
	for i := 0; i < b.N; i++ {
		id := fmt.Sprintf("User%d", i)
		player = contracts.User{Id: id, Name: id}
		matching.AddToGroup(player)
	}
	b.Logf("%d groups have been formed.", len(matching.Groups))
}

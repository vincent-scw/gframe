package game

import (
	"math/rand"
	"time"

	e "github.com/vincent-scw/gframe/kafka/events"
)

type groupStatus int

const (
	forming groupStatus = iota
	formed  groupStatus = iota
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Group represents a group of players
type Group struct {
	ID      string
	Players []*e.User
	status  groupStatus
}

// Matching represents a matching game
type Matching struct {
	Groups    map[string]Group
	GroupSize int
}

// NewMatching returns Matching
func NewMatching() *Matching {
	matching := Matching{}
	// a group has 2 players, 1v1.
	// TODO: configuration
	matching.GroupSize = 2
	return &matching
}

// AddToGroup adds a player to group
func (m *Matching) AddToGroup(p *e.User) {
	if m.Groups == nil {
		m.Groups = make(map[string]Group, m.GroupSize)
	}
	group := m.findOrCreateFormingGroup()
	group.Players = append(group.Players, p)
	if len(group.Players) == m.GroupSize {
		group.status = formed
	}
}

func (m *Matching) findOrCreateFormingGroup() *Group {
	for _, group := range m.Groups {
		if group.status == forming {
			return &group
		}
	}

	id := randSeq(7)
	g := Group{ID: id, status: forming}
	m.Groups[id] = g
	return &g
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

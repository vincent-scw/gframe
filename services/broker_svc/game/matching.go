package game

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/vincent-scw/gframe/broker_svc/singleton"
	e "github.com/vincent-scw/gframe/contracts"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Group represents a group of players
type Group struct {
	e.GroupInfo
	userChan chan e.User
	killChan chan struct{}
}

// Matching represents a matching game
type Matching struct {
	Groups    map[string]*Group
	GroupSize int
	// FormingTimeout by seconds
	FormingTimeout int
	util
}

type util struct {
	formingGroup *Group
	confirmChan  chan bool
}

// NewMatching returns Matching
func NewMatching(groupSize int, maxGroupCount int, timeoutInSeconds int) *Matching {
	matching := Matching{GroupSize: groupSize, FormingTimeout: timeoutInSeconds}

	matching.Groups = make(map[string]*Group, maxGroupCount)
	matching.confirmChan = make(chan bool)

	return &matching
}

// AddToGroup adds a player to group
func (m *Matching) AddToGroup(player e.User) bool {
	m.prepareFormingGroup()
	if m.formingGroup == nil {
		return false
	}
	m.formingGroup.userChan <- player
	return <-m.confirmChan
}

func (m *Matching) prepareFormingGroup() {
	if m.formingGroup == nil {
		m.formingGroup = newGroup(m.GroupSize)
		go m.formingGroup.formGroup(m)
	}
}

func (m *Matching) groupFormed() {
	g := m.formingGroup
	if g != nil && len(g.Players) > 1 {
		g.Status = e.GroupFormed
		m.Groups[g.ID] = g

		value, _ := json.Marshal(g)
		go singleton.RedisPublish(e.GroupChannel, string(value))
	}
	m.formingGroup = nil
}

func newGroup(groupSize int) *Group {
	id := randSeq(7)
	g := Group{GroupInfo: e.GroupInfo{ID: id, Status: e.GroupForming}}
	g.userChan = make(chan e.User)
	g.killChan = make(chan struct{})

	return &g
}

func (g *Group) formGroup(m *Matching) {
	t := time.After(time.Second * time.Duration(m.FormingTimeout))
	for {
		select {
		case <-t:
			// Timeout
			m.groupFormed()
		case u := <-g.userChan:
			if len(g.Players) < m.GroupSize {
				g.Players = append(g.Players, u)
				if len(g.Players) == m.GroupSize {
					m.groupFormed()
				}
				m.confirmChan <- true
			} else {
				m.confirmChan <- false
			}
		}
	}
}

func (g *Group) closeChan() {
	close(g.userChan)
	close(g.killChan)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

package game

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
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
	lock         sync.RWMutex
	formingGroup *Group
	pubsubClient *r.PubSubClient
}

// NewMatching returns Matching
func NewMatching(groupSize int, maxGroupCount int, timeoutInSeconds int) *Matching {
	matching := Matching{GroupSize: groupSize, FormingTimeout: timeoutInSeconds}

	matching.lock = sync.RWMutex{}
	matching.Groups = make(map[string]*Group, maxGroupCount)
	matching.pubsubClient = r.NewPubSubClient("40.83.112.48:6379")

	return &matching
}

// AddToGroup adds a player to group
func (m *Matching) AddToGroup(player e.User) bool {
	time.Sleep(time.Millisecond * time.Duration(20)) // Wait for last forming group to complete
	m.prepareFormingGroup()

	m.formingGroup.userChan <- player
	return true
}

func (m *Matching) prepareFormingGroup() {
	if m.formingGroup == nil {
		m.lock.Lock()
		defer m.lock.Unlock()
		if m.formingGroup == nil {
			m.formingGroup = newGroup(m.GroupSize)
			go m.formingGroup.formGroup(m)
			go m.waitForKill()
		}
	}
}

func (m *Matching) waitForKill() {
	<-m.formingGroup.killChan

	m.lock.Lock()
	defer m.lock.Unlock()
	g := m.formingGroup
	if len(g.Players) > 1 {
		g.Status = e.GroupFormed
		m.Groups[g.ID] = g
		value, _ := json.Marshal(g)
		go m.pubsubClient.Publish(e.GroupChannel, string(value))
	}
	m.formingGroup = nil
}

// Close releases resources
func (m *Matching) Close() {
	m.pubsubClient.Close()
}

func newGroup(groupSize int) *Group {
	id := randSeq(7)
	g := Group{GroupInfo: e.GroupInfo{ID: id, Status: e.GroupForming}}
	g.userChan = make(chan e.User, groupSize)
	g.killChan = make(chan struct{})
	return &g
}

func (g *Group) formGroup(m *Matching) {
	t := time.After(time.Second * time.Duration(m.FormingTimeout))
	for {
		select {
		case <-t:
			// Timeout
			g.killChan <- struct{}{}
		case u := <-g.userChan:
			g.Players = append(g.Players, u)
			if len(g.Players) == m.GroupSize {
				g.killChan <- struct{}{}
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

package game

import (
	"math/rand"
	"sync"
	"time"

	e "github.com/vincent-scw/gframe/kafkactl/events"
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
	ID       string
	Players  []e.User
	status   groupStatus
	userChan chan e.User
	killChan chan struct{}
}

// Matching represents a matching game
type Matching struct {
	Groups    map[string]*Group
	GroupSize int
	// FormingTimeout by seconds
	FormingTimeout int
	lock           sync.RWMutex
	formingGroup   *Group
}

// NewMatching returns Matching
func NewMatching(groupSize int, maxGroupCount int, timeoutInSeconds int) *Matching {
	matching := Matching{GroupSize: groupSize, FormingTimeout: timeoutInSeconds,
		lock: sync.RWMutex{}}

	matching.Groups = make(map[string]*Group, maxGroupCount)

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
		g.status = formed
		m.Groups[g.ID] = g
	}
	m.formingGroup = nil
}

func newGroup(groupSize int) *Group {
	id := randSeq(7)
	g := Group{ID: id, status: forming}
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

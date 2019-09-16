package game

import (
	"time"

	c "github.com/vincent-scw/gframe/contracts"
	u "github.com/vincent-scw/gframe/util"
)

// GroupFormed event
type GroupFormed func(formed *Group)

// Group represents a group of players
type Group struct {
	c.GroupInfo
	userChan chan c.User
	killChan chan struct{}
}

// Matching represents a matching game
type Matching struct {
	Groups    map[string]*Group
	GroupSize int
	// FormingTimeout by seconds
	FormingTimeout int
	Formed         GroupFormed
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
func (m *Matching) AddToGroup(player c.User) bool {
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
		g.Status = c.GroupStatus_Formed
		m.Groups[g.Id] = g
		// Event out
		if m.Formed != nil {
			m.Formed(g)
		}
	}
	m.formingGroup = nil
}

func newGroup(groupSize int) *Group {
	id := u.NextRandom()
	g := Group{GroupInfo: c.GroupInfo{Id: id, Status: c.GroupStatus_Forming}}
	g.userChan = make(chan c.User)
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
				g.Players = append(g.Players, &u)
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

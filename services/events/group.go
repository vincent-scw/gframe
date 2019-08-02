package events

// GroupStatus is enum
type GroupStatus int

const (
	// GroupForming forming
	GroupForming GroupStatus = 201
	// GroupFormed formed
	GroupFormed GroupStatus = 202
	// GroupChannel Redis channel
	GroupChannel string = "group_chan"
)

// GroupInfo represents basic info
type GroupInfo struct {
	ID      string      `json:"id"`
	Players []User      `json:"players"`
	Status  GroupStatus `json:"status"`
}

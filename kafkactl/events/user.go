package events

// Status is enum
type Status int

const (
	// EventOut out
	EventOut Status = iota
	// EventIn in
	EventIn
)

// User is a model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserEvent is a model with event
type UserEvent struct {
	User
	Type Status `json:"type"`
}

// DefKey implements KeyDef interface
func (userEvent *UserEvent) DefKey() string {
	return userEvent.ID
}

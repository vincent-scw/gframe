package events

// UserStatus is enum
type UserStatus int

const (
	// UserEventOut out
	UserEventOut UserStatus = 101
	// UserEventIn in
	UserEventIn UserStatus = 102
)

// User is a model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserEvent is a model with event
type UserEvent struct {
	User
	Type UserStatus `json:"type"`
}

// DefKey implements KeyDef interface
func (userEvent *UserEvent) DefKey() string {
	return userEvent.ID
}

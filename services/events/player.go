package events

// UserStatus is enum
type UserStatus int

const (
	// UserEventOut out
	UserEventOut UserStatus = 101
	// UserEventIn in
	UserEventIn UserStatus = 102
	// PlayerChannel Redis channel
	PlayerChannel string = "player_chann"
)

// PlayerEvent is a model with event
type PlayerEvent struct {
	User
	Type UserStatus `json:"type"`
}

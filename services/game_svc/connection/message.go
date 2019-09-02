package connection

// EventType represent message type
type EventType string

const (
	// Group message
	Group EventType = "group"
	// Player message
	Player EventType = "player"
	// Game message
	Game EventType = "game"
)

// Message struct
type Message struct {
	Type    EventType
	Content []byte
}

// NewMessage returns a new message
func NewMessage(t EventType, content []byte) *Message {
	return &Message{Type: t, Content: content}
}

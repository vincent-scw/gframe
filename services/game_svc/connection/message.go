package connection

// EventType represent message type
type EventType string

const (
	// Common message
	Common EventType = "common"
	// Gaming message
	Gaming EventType = "gaming"
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

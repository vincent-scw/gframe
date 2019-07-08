package event

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

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

func getUserFromToken(token *jwt.Token) (*User, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(string)
		// sid
		return &User{ID: sub, Name: sub}, nil
	}

	return nil, errors.New("cannot read info from token")
}

// NewEvent generates an UserEvent based on givent token and event type
func NewEvent(token *jwt.Token, t Status) (userEvent *UserEvent, err error) {
	user, err := getUserFromToken(token)
	if err == nil {
		userEvent = &UserEvent{User: *user, Type: t}
	}
	return
}

// DefKey implements KeyDef interface
func (userEvent *UserEvent) DefKey() string {
	return userEvent.ID
}

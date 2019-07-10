package event

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	e "github.com/vincent-scw/gframe/kafka/events"
)

func getUserFromToken(token *jwt.Token) (*e.User, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(string)
		// sid
		return &e.User{ID: sub, Name: sub}, nil
	}

	return nil, errors.New("cannot read info from token")
}

// NewEvent generates an UserEvent based on givent token and event type
func NewEvent(token *jwt.Token, t e.Status) (userEvent *e.UserEvent, err error) {
	user, err := getUserFromToken(token)
	if err == nil {
		userEvent = &e.UserEvent{User: *user, Type: t}
	}
	return
}

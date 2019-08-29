package auth

import (
	"github.com/rs/xid"
)

// Player entity
type Player struct {
	Name string `json:"name"`
}

// AccessToken entity
type AccessToken struct {
	Token string `json:"token"`
}

// ToToken to token
func (p *Player) ToToken() *AccessToken {
	guid := xid.New()
	token, _ := generateJwtToken(guid.String(), p.Name)
	return &AccessToken{Token: token}
}


package auth

import (
	"github.com/rs/xid"
)

// Player entity
type Player struct {
	Name string `json:"name"`
}

// Token entity
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// ToToken to token
func (p *Player) ToToken() *Token {
	guid := xid.New()
	token, _ := generateJwtToken(guid.String(), p.Name)
	return &Token{AccessToken: token, TokenType: "Bearer"}
}


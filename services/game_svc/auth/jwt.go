package auth

import (
	"errors"

	"github.com/kataras/iris"
	"github.com/dgrijalva/jwt-go"
	jwtmid "github.com/iris-contrib/middleware/jwt"

	"github.com/vincent-scw/gframe/events"
	"github.com/vincent-scw/gframe/game_svc/config"
)

// JwtHandler handles jwt token
var JwtHandler = jwtmid.New(jwtmid.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtKey()), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// GetUserFromToken reads user info from token
func GetUserFromToken(ctx iris.Context, status events.User_Status) (*events.User, error) {
	authToken := JwtHandler.Get(ctx)
	if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		sub := claims["sub"].(string)
		// sid
		return &events.User{Id: sub, Name: sub, Status: status}, nil
	}

	return nil, errors.New("cannot read info from token")
}
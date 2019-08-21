package auth

import (
	"errors"

	"github.com/kataras/iris"
	"github.com/dgrijalva/jwt-go"
	jwtmid "github.com/iris-contrib/middleware/jwt"

	"github.com/vincent-scw/gframe/contracts"
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
func GetUserFromToken(ctx iris.Context, status contracts.User_Status) (*contracts.User, error) {
	authToken := JwtHandler.Get(ctx)
	if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		sub := claims["sub"].(string)
		// sid
		return &contracts.User{Id: sub, Name: sub, Status: status}, nil
	}

	return nil, errors.New("cannot read info from token")
}
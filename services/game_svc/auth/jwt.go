package auth

import (
	"errors"
	"log"
	"time"

	"github.com/kataras/iris"
	"github.com/dgrijalva/jwt-go"
	jwtmid "github.com/iris-contrib/middleware/jwt"

	"github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/game_svc/config"
)

// JwtHandler handles jwt token
var JwtHandler = jwtmid.New(jwtmid.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return config.GetJwtKey(), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// WSJwtHandler handles jwt token for websocket
var WSJwtHandler = jwtmid.New(jwtmid.Config{
	Extractor: jwtmid.FromParameter("token"),
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return config.GetJwtKey(), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// GetUserFromToken reads user info from token
func GetUserFromToken(ctx iris.Context, status contracts.User_Status) (*contracts.User, error) {
	authToken := JwtHandler.Get(ctx)
	return toUser(authToken, status)
}

// GetUserFromTokenForWS reads user info from token for websocket
func GetUserFromTokenForWS(ctx iris.Context, status contracts.User_Status) (*contracts.User, error) {
	authToken := WSJwtHandler.Get(ctx)
	return toUser(authToken, status)
}

func toUser(authToken *jwt.Token, status contracts.User_Status) (*contracts.User, error) {
	if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		id := claims["id"].(string)
		sub := claims["sub"].(string)
		// sid
		return &contracts.User{Id: id, Name: sub, Status: status}, nil
	}

	return nil, errors.New("cannot read info from token")
}

func generateJwtToken(id, name string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iss"] = "gframe_game"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["sub"] = id
	claims["name"] = name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetJwtKey())

	if err != nil {
		log.Printf("Error signing token: %v\n", err)
	}
	return tokenString, err
}
package auth

import (
	"fmt"
	"testing"
	"github.com/dgrijalva/jwt-go"
)

func TestGenToken(t *testing.T) {
	tokenString, _ := generateJwtToken("aaa", "test")
	token, err := jwt.Parse(tokenString, func(j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse error")
		}
		return []byte("00000000"), nil
	})

	if err != nil {

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"].(string) != "aaa" || claims["name"].(string) != "test" {
			t.Error("claims generating failed.")	
		}
	} else {
		t.Error("token generating failed.")
	}
}
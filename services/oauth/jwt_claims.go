package main

// import (
// 	"time"
// 	"errors"
	
// 	"github.com/dgrijalva/jwt-go"
// )

// // JWTAccessClaims claims
// type JWTAccessClaims struct {
// 	jwt.StandardClaims

// 	Custom string `json:"cust,omitempty"`
// }

// // Valid claims verification
// func (a *JWTAccessClaims) Valid() error {
// 	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
// 		return errors.ErrInvalidAccessToken
// 	}
// 	return nil
// }
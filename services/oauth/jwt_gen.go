package main

// import (
// 	"github.com/dgrijalva/jwt-go"
// 	"gopkg.in/oauth2.v3"
// 	"gopkg.in/oauth2.v3/errors"
// 	"gopkg.in/oauth2.v3/utils/uuid"
// )

// // JWTAccessGenerate generate the jwt access token
// type JWTAccessGenerate struct {
// 	SignedKey    []byte
// 	SignedMethod jwt.SigningMethod
// }

// // NewJWTAccessGenerate create to generate the jwt access token instance
// func NewJWTAccessGenerate(key []byte, method jwt.SigningMethod) *JWTAccessGenerate {
// 	return &JWTAccessGenerate{
// 		SignedKey:    key,
// 		SignedMethod: method,
// 	}
// }

// func (a *JWTGenerator) Token(data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
// 	claims := &JWTAccessClaims{
// 		StandardClaims: jwt.StandardClaims{
// 			Audience:  data.Client.GetID(),
// 			Subject:   data.UserID,
// 			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(a.SignedMethod, claims)
// 	var key interface{}
// 	if a.isEs() {
// 		key, err = jwt.ParseECPrivateKeyFromPEM(a.SignedKey)
// 		if err != nil {
// 			return "", "", err
// 		}
// 	} else if a.isRsOrPS() {
// 		key, err = jwt.ParseRSAPrivateKeyFromPEM(a.SignedKey)
// 		if err != nil {
// 			return "", "", err
// 		}
// 	} else if a.isHs() {
// 		key = a.SignedKey
// 	} else {
// 		return "", "", errs.New("unsupported sign method")
// 	}
// 	access, err = token.SignedString(key)
// 	if err != nil {
// 		return
// 	}

// 	if isGenRefresh {
// 		refresh = base64.URLEncoding.EncodeToString(uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).Bytes())
// 		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
// 	}

// 	return
// }

// func (a *JWTAccessGenerate) isEs() bool {
// 	return strings.HasPrefix(a.SignedMethod.Alg(), "ES")
// }

// func (a *JWTAccessGenerate) isRsOrPS() bool {
// 	isRs := strings.HasPrefix(a.SignedMethod.Alg(), "RS")
// 	isPs := strings.HasPrefix(a.SignedMethod.Alg(), "PS")
// 	return isRs || isPs
// }

// func (a *JWTAccessGenerate) isHs() bool {
// 	return strings.HasPrefix(a.SignedMethod.Alg(), "HS")
// }
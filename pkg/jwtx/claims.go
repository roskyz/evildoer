package jwtx

import (
	"encoding/json"
	"evildoer/pkg/cryptic"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SetDefaultIssuer is must be called
func SetDefaultIssuer(issuer Issuer) {
	defaultIssuer = issuer
}

type Claims struct {
	Values string
	jwt.RegisteredClaims
}

type credentials interface {
	Claims(secret cryptic.Secret) (string, error)
	Values(values string, secret cryptic.Secret) error
}

func Signin(credentials credentials, expireTime int) (token string, err error) {
	return defaultIssuer.Signin(credentials, expireTime)
}

func Welcome(tknStr string, credentials credentials) error {
	return defaultIssuer.Welcome(tknStr, credentials)
}

func GetExpireTime(tknStr string) (expiredAt time.Time) {
	token, _, err := new(jwt.Parser).ParseUnverified(tknStr, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("todo")
	}
	switch exp := claims["exp"].(type) {
	case float64:
		expiredAt = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		expiredAt = time.Unix(v, 0)
	}
	return expiredAt
}

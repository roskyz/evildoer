package jwtx

import (
	"evildoer/pkg/cryptic"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var defaultIssuer Issuer

type Issuer struct {
	signingMethod jwt.SigningMethod
	timeBase      time.Duration
	secret        cryptic.Secret
}

func NewIssuer(signingMethod jwt.SigningMethod, timeBase time.Duration, secret cryptic.Secret) Issuer {
	return Issuer{
		signingMethod: signingMethod, timeBase: timeBase, secret: secret,
	}
}

func (i *Issuer) Signin(credentials credentials, expireTime int) (token string, err error) {
	values, err := credentials.Claims(i.secret)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(time.Duration(expireTime) * i.timeBase)
	claims := &Claims{
		Values: values,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	// Create the JWT string
	if token, err = jwt.NewWithClaims(i.signingMethod, claims).SignedString([]byte(i.secret)); err != nil {
		return "", err
	}
	return token, nil
}

func (i *Issuer) Welcome(tknStr string, credentials credentials) error {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(i.secret), nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		return jwt.ErrTokenUnverifiable
	}
	return credentials.Values(claims.Values, i.secret)
}

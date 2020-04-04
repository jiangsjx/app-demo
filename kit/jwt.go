package kit

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}

type AuthJWT struct {
	SigningKey []byte
}

const SigningKey = "12345678"

func NewAuthJWT() *AuthJWT {
	return &AuthJWT{SigningKey: []byte(SigningKey)}
}

func (auth *AuthJWT) CreateToken(claims *AuthClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(auth.SigningKey)
}

func (auth *AuthJWT) ParseToken(tokenStr string) (*AuthClaims, error) {
	keyFunc := func(token *jwt.Token) (i interface{}, err error) {
		return auth.SigningKey, nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, keyFunc)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

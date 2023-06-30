package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JwtClaims struct {
	ID         uint  `json:"sub,omitempty"`
	ExpieresAt int64 `json:"exp,omitempty"`
	IssuedAt   int64 `json:"iat,omitempty"`
}

func (c JwtClaims) Valid(helper *jwt.ValidationHelper) (err error) {
	if helper.After(time.Unix(c.ExpieresAt, 0)) {
		err = errors.New("token has expired")
	}
	if helper.Before(time.Unix(c.IssuedAt, 0)) {
		err = errors.New("token used before issued")
	}
	return err
}

func GenerateToken(user User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 4)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{ID: user.ID, ExpieresAt: expiry.Unix(), IssuedAt: now.Unix()})
	tokenString, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		panic(err)
	}

	return tokenString, err
}

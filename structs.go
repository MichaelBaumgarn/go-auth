package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type User struct {
	ID       uint   `gorm:"column:user_id" gorm:"primaryKey" `
	Email    string `gorm:"unique,column:email" json:"email" binding:"required"`
	Password string `gorm:"column:password" json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Response struct {
	Code  int         `json:"code,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
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

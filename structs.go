package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"gorm.io/gorm"
)

type User struct {
	db       *gorm.DB
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

var DB *gorm.DB

func (c JwtClaims) Valid(helper *jwt.ValidationHelper) (err error) {
	if helper.After(time.Unix(c.ExpieresAt, 0)) {
		err = errors.New("token has expired")
	}
	if helper.Before(time.Unix(c.IssuedAt, 0)) {
		err = errors.New("token used before issued")
	}
	return err
}

func (u User) GetAll() []User {
	var users []User
	u.db.Find(&users)
	return users
}

func (u User) GetByID(id string) User {
	fmt.Printf("see id %v", id)
	var user User
	u.db.Find(&user, id)
	return user
}

func (u User) Create(user User) User {
	u.db.Create(&user)
	return user
}

func (u User) Authenticate(incomingUser User) (string, error) {
	var user User
	u.db.First(&user, "email = ?", &incomingUser.Email)

	fmt.Printf("%v %v", user.Password, incomingUser.Password)
	fmt.Println("foobar")
	fmt.Printf("user %v", user.Password)
	fmt.Printf("user %v", incomingUser.Password)
	if user.Password != incomingUser.Password {
		return "", errors.New("wrong password")
	}

	var token string
	var err error
	if token, err = generateToken(user); err != nil {
		return "", errors.New("could not generate token")
	}
	return token, err
}

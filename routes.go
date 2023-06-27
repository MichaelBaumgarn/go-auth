package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}

func (config Config) GetUser(c *gin.Context) {
	var user User
	config.db.Find(&user, 1)
	c.IndentedJSON(http.StatusOK, user)
}

func (config Config) PostUser(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}
	config.db.Create(&incomingUser)

	fmt.Printf("%v", &incomingUser)

	var users []User
	config.db.Model(&User{}).Find(&users)
	c.IndentedJSON(http.StatusOK, &users)
}

func (config Config) Login(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}

	var user User
	config.db.First(&user, "email = ?", &incomingUser.Email)

	fmt.Printf("%v %v", user.Password, incomingUser.Password)
	fmt.Println("foobar")
	fmt.Printf("user %v", user.Password)
	fmt.Printf("user %v", incomingUser.Password)
	if user.Password != incomingUser.Password {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: "bad password"})
	}

	var token string
	var err error
	if token, err = generateToken(user); err != nil {
		c.JSON(http.StatusUnauthorized, Response{Code: 400, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Data: fmt.Sprintf("Bearer %s", token)})
}

func (config Config) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var user User
	config.db.Find(&user, id)
	c.IndentedJSON(http.StatusOK, user)
}

func generateToken(user User) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 4)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, JwtClaims{ID: user.ID, ExpieresAt: expiry.Unix(), IssuedAt: now.Unix()})
	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return tokenString, err
}

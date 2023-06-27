package main

import (
	"fmt"
	"net/http"

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
	if user.Password == incomingUser.Password {
		c.IndentedJSON(http.StatusOK, &user)
	} else {

		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: "bad password"})
	}
}

func (config Config) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var user User
	config.db.Find(&user, id)
	c.IndentedJSON(http.StatusOK, user)
}

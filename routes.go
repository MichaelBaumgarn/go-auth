package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e Env) GetUser(c *gin.Context) {
	users := e.users.GetAll()

	c.IndentedJSON(http.StatusOK, users)
}

func (e Env) PostUser(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}
	e.users.Create(incomingUser)

	fmt.Printf("%v", &incomingUser)

	c.IndentedJSON(http.StatusOK, &incomingUser)
}

func (e Env) Login(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}

	var token string
	var err error
	if token, err = e.users.Authenticate(incomingUser); err != nil {
		c.JSON(http.StatusUnauthorized, Response{Code: 400, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Data: fmt.Sprintf("Bearer %s", token)})
}

func (e Env) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// user := UserModel{db:}.ModelGetUserByID(id, config)
	user := e.users.GetByID(id)

	c.IndentedJSON(http.StatusOK, user)
}

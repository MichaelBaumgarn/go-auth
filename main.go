package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// album represents data about a record album.
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

type config struct {
	db *gorm.DB
}

func main() {
	dsn := "host=localhost dbname=vocab port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
		panic("db connection failed")
	}
	c := config{db: db}

	router := gin.Default()
	router.GET("/user", c.getUser)
	router.GET("/user/:id", c.getUserByID)
	router.POST("/user", c.postUser)

	router.Run("localhost:8081")
}

func (config config) getUser(c *gin.Context) {
	var user User
	config.db.Find(&user, 1)
	c.IndentedJSON(http.StatusOK, user)
}

func (config config) postUser(c *gin.Context) {
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

func (config config) getUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var user User
	config.db.Find(&user, id)
	c.IndentedJSON(http.StatusOK, user)
}

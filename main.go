package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// album represents data about a record album.
func main() {
	dsn := "host=localhost dbname=vocab port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
		panic("db connection failed")
	}
	c := Config{db: db}

	router := gin.Default()
	router.GET("/user", c.GetUser)
	router.GET("/user/:id", c.GetUserByID)
	router.POST("/user", c.PostUser)

	router.POST("/login", c.Login)

	router.Run("localhost:8081")
}

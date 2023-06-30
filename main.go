package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	users User
}

func main() {
	dsn := "host=localhost dbname=vocab port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
		panic("db connection failed")
	}

	env := &Env{
		users: User{db: db},
	}

	router := gin.Default()
	router.GET("/user", env.GetUser)
	router.GET("/user/:id", env.GetUserByID)
	router.POST("/user", env.PostUser)

	router.POST("/login", env.Login)

	router.Run("localhost:8081")
}

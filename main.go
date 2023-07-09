package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	users       User
	grammar     Grammar
	userGrammar UserGrammar
}

func SetupRouter(env Env) *gin.Engine {
	router := gin.Default()
	router.GET("/user", env.GetUser)
	router.GET("/user/:id", env.GetUserByID)
	router.POST("/user", env.PostUser)

	router.POST("/login", env.Login)

	router.GET("/grammar", env.GetAllGrammar)

	router.Run("localhost:8081")
	return router
}

func main() {
	dsn := "host=localhost dbname=vocab port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
		panic("db connection failed")
	}

	SetupRouter(Env{
		users:       User{db: db},
		grammar:     Grammar{db: db},
		userGrammar: UserGrammar{db: db},
	})
	// router.Use(
	// 	AuthMiddleware,
	// )

}

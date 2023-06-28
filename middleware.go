package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		// not auth
		c.Next()
	}
	// claims, err := parseToken(token)
}

func ParseToken(tokenString, secret string) (claims JwtClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return JwtClaims{}, fmt.Errorf("invalid token %v", err)
	}
	return
}

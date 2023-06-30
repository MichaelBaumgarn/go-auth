package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.Set(IsAuthenticatedKey, false)
		c.Next()
		return
	}
	claims, err := ParseToken(token, TokenSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{Code: 400, Error: err.Error()})
	}
	c.Set(IsAuthenticatedKey, true)
	c.Set(UserIDKey, claims.ID)
	c.Next()
}

func ParseToken(tokenString, secret string) (claims JwtClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return JwtClaims{}, fmt.Errorf("invalid token %v", err)
	}
	return
}

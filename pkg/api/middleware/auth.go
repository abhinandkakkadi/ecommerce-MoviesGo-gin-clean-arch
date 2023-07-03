package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(c *gin.Context) {

	s := c.Request.Header.Get("Authorization")

	var token string
	if s[:7] == "Bearer " {
		token = strings.TrimPrefix(s, "Bearer ")
	} else {
		token = s
	}

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

func validateToken(token string) error {

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("12345678"), nil
	})

	return err

}

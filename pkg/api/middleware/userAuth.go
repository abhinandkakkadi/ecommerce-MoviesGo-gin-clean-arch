package middleware

import (
	"net/http"

	helper "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		tokenString := helper.GetTokenFromHeader(authHeader)
		// Validate the token and extract the user ID
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Add the user ID to the Gin context
		c.Set("user_id", userID)
		c.Set("user_email", userEmail)

		// Call the next handler
		c.Next()
	}
}

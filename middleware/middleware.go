package middleware

import (
	"net/http"
	"strings"

	"github.com/Atipat-CMU/api-gateway/internal/service"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookies
		token, err := c.Cookie("Authorization")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if it exists
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate the token (you can implement your own validation logic here)
		userId, err := service.NewAuthService().ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// You can store the user ID in the context to access later in the request handlers
		c.Request.Header.Set("X-User-Id", userId)

		// Proceed to the next handler
		c.Next()
	}
}

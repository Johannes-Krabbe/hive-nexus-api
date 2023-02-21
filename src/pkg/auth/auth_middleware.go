package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a middleware function to verify JWT tokens
func VerifyTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token string from the Authorization cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Authorization token cookie is missing"})
			return
		}

		// Verify the token and get the user ID
		userID, err := verifyJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Invalid token"})
			return
		}

		// Add the user ID to the request context for later use
		c.Set("UserID", userID)

		c.Next()
	}
}

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Auth Token!"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateAuthToken(tokenString)

		// Add token claims to context
		c.Set("currentUserEmail", claims["email"])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}

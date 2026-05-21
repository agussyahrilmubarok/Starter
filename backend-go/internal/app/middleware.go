package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (app *App) requireAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing authorization header",
			})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			return
		}

		tokenString := splitToken[1]

		userID, err := app.jwtService.Validate(c.Request.Context(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}

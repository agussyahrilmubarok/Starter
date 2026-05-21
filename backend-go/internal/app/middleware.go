package app

import (
	"net/http"
	"strings"

	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (app *App) requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		log := logger.Get().With(zap.String("request_id", requestID))
		ctx := logger.WithCtx(c.Request.Context(), log)

		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}

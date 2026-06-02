package middleware

import (
	"net/http"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/delivery/http/payload"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"github.com/gin-gonic/gin"
)

func Auth(jwtManager security.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, payload.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"error": "Authorization header is required"},
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, payload.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"error": "Invalid authorization format, expected: Bearer {token}"},
			})
			return
		}

		userID, err := jwtManager.ValidateToken(parts[1])
		if err != nil {
			msg := "Invalid token"
			if err == security.ErrExpiredToken {
				msg = "Token has expired"
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, payload.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"error": msg},
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

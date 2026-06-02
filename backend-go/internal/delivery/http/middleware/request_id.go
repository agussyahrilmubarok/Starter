package middleware

import (
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestIDMiddleware() gin.HandlerFunc {
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

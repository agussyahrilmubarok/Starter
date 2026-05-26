package application

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/controller"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

func (app *App) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		user, err := controller.GetUserSession(session)
		if err != nil {
			if errors.Is(err, controller.ErrSessionNotFound) {
				c.Redirect(http.StatusFound, "/sign-in")
				c.Abort()
				return
			}

			controller.ClearUserSession(session)
			c.Redirect(http.StatusFound, "/sign-in")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (app *App) guestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		_, err := controller.GetUserSession(session)
		if err == nil {
			c.Redirect(http.StatusFound, "/dashboard")
			c.Abort()
			return
		}

		c.Next()
	}
}

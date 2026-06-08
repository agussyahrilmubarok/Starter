package middleware

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		user, err := session.GetUser(s)
		if err != nil {
			if errors.Is(err, session.ErrSessionNotFound) {
				c.Redirect(http.StatusFound, "/sign-in")
				c.Abort()
				return
			}

			session.DeleteUser(s)
			c.Redirect(http.StatusFound, "/sign-in")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

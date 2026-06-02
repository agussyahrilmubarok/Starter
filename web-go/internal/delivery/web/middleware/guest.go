package middleware

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GuestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)

		_, err := session.GetUser(s)
		if err == nil {
			c.Redirect(http.StatusFound, "/dashboard")
			c.Abort()
			return
		}

		c.Next()
	}
}

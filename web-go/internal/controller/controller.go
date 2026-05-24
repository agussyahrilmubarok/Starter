package controller

import (
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, code int, tmpl string, data gin.H) {
	c.HTML(code, tmpl, data)
}

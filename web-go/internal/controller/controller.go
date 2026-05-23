package controller

import (
	"github.com/gin-gonic/gin"
)

type templateData struct {
	Title  string            // document.title
	Values map[string]string // .Values.Name = "John Doe"
	Errors map[string]string // .Errors.Email = "Email is not found"
}

func render(c *gin.Context, code int, tmpl string, data templateData) {
	c.HTML(code, tmpl, data)
}

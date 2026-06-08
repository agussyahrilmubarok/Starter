package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AppController) HomePage(c *gin.Context) {
	data := gin.H{
		"Title": "Home",
	}

	render(c, http.StatusOK, "home_index.html", data)
}

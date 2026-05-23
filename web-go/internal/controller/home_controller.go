package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type homeController struct {
}

func (h *homeController) Index(c *gin.Context) {
	data := templateData{
		Title: "Home",
	}

	render(c, http.StatusOK, "home_index.html", data)
}

func NewHomeController() *homeController {
	return &homeController{}
}

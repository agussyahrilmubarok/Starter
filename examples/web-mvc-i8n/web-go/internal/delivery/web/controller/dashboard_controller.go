package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AppController) DashboardPage(c *gin.Context) {
	data := gin.H{
		"Title": "Dashboard",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	render(c, http.StatusOK, "dashboard_index.html", data)
}

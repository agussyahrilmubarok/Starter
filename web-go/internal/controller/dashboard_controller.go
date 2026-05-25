package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dashboardController struct {
}

func (h *dashboardController) Dashboard(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	data := gin.H{
		"Title":       "Dashboard",
		"UserProfile": userProfile,
	}

	render(c, http.StatusOK, "dashboard_index.html", data)
}

func NewDashboardController() *dashboardController {
	return &dashboardController{}
}

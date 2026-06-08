package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
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

	s := sessions.Default(c)
	if flashes := s.Flashes("success"); len(flashes) > 0 {
		data["MsgSuccess"] = flashes[0]
	}
	if flashes := s.Flashes("error"); len(flashes) > 0 {
		data["MsgError"] = flashes[0]
	}
	s.Save()

	render(c, http.StatusOK, "dashboard_index.html", data)
}

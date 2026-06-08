package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *AppController) ProfilePage(c *gin.Context) {
	data := gin.H{
		"Title": "Profile",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	words := strings.Fields(userProfile.Name)
	initials := ""
	for _, word := range words {
		if len(word) > 0 {
			initials += string([]rune(word)[0])
		}
	}
	data["UserProfileImg"] = strings.ToUpper(initials)

	render(c, http.StatusOK, "profile_index.html", data)
}

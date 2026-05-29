package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type profileController struct {
}

func (h *profileController) Index(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	data := gin.H{
		"Title":          "Profile",
		"UserProfile":    userProfile,
		"UserProfileImg": h.getInitials(userProfile.Name),
	}
	data["UserProfileImg"] = string([]rune(userProfile.Name)[0])

	render(c, http.StatusOK, "profile_index.html", data)
}

func NewProfileController() *profileController {
	return &profileController{}
}

func (h *profileController) getInitials(name string) string {
	words := strings.Fields(name)
	initials := ""
	for _, word := range words {
		if len(word) > 0 {
			initials += string([]rune(word)[0])
		}
	}
	return strings.ToUpper(initials)
}

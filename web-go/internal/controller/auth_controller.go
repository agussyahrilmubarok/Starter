package controller

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/model"
	"github.com/gin-gonic/gin"
)

type authController struct {
}

func (h *authController) SignUpPage(c *gin.Context) {
	data := templateData{
		Title: "Sign Up",
	}

	render(c, http.StatusOK, "sign_up_index.html", data)
}

func (h *authController) SignUp(c *gin.Context) {
	data := templateData{
		Title: "Sign Up",
	}

	var req model.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		data.Errors[""] = 
		render(c, http.StatusBadRequest, "sign_up_index.html", data)
		return
	}

	render(c, http.StatusOK, "sign_up_index.html", data)
}

func NewAuthController() *authController {
	return &authController{}
}

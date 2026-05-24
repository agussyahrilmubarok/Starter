package controller

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/model"
	"agussyahrilmubarok.github.io/web/internal/service"
	"agussyahrilmubarok.github.io/web/pkg/helper"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService service.IAuthService
}

func (h *authController) SignUp(c *gin.Context) {
	data := gin.H{
		"Title": "Sign Up",
	}

	var req model.SignUpRequest
	if c.Request.Method == http.MethodGet {
		data["Values"] = req
		render(c, http.StatusOK, "sign_up_index.html", data)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		data["Values"] = req
		data["Message"] = "Validation error"
		data["Errors"] = helper.ValidatorError(err)
		render(c, http.StatusBadRequest, "sign_up_index.html", data)
		return
	}

	_, err := h.authService.SignUp(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailExists:
			errorsMap["Email"] = "Email already exists"
			data["Values"] = req
			data["Message"] = "Conflict error"
			data["Errors"] = errorsMap
			render(c, http.StatusBadRequest, "sign_up_index.html", data)
		default:
			errorsMap["Error"] = "Something went wrong"
			data["Values"] = req
			data["Message"] = "Internal server error"
			data["Errors"] = errorsMap
			render(c, http.StatusBadRequest, "sign_up_index.html", data)
		}
		return
	}

	data["MsgInfo"] = "Sign up successfully"
	render(c, http.StatusBadRequest, "sign_up_index.html", data)
}

func NewAuthController(
	authService service.IAuthService,
) *authController {
	return &authController{
		authService: authService,
	}
}

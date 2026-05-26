package controller

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/model"
	"agussyahrilmubarok.github.io/web/internal/service"
	"agussyahrilmubarok.github.io/web/pkg/validatorutil"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService service.IAuthService
}

// Index — GET/POST /sign-up
func (h *authController) SignUp(c *gin.Context) {
	session := sessions.Default(c)

	data := gin.H{
		"Title": "Sign Up",
	}

	if flashes := session.Flashes("success"); len(flashes) > 0 {
		data["MsgInfo"] = flashes[0]
	}
	session.Save()

	var req model.SignUpRequest
	if c.Request.Method == http.MethodGet {
		data["Values"] = req
		render(c, http.StatusOK, "sign_up_index.html", data)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		data["Values"] = req
		data["Message"] = "Validation error"
		data["Errors"] = validatorutil.ValidatorError(err)
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

	session.AddFlash("Sign up successfully!", "success")
	session.Save()

	c.Redirect(http.StatusFound, "/sign-in")
}

// Index — GET/POST /sign-in
func (h *authController) SignIn(c *gin.Context) {
	session := sessions.Default(c)

	data := gin.H{
		"Title": "Sign In",
	}

	if flashes := session.Flashes("success"); len(flashes) > 0 {
		data["MsgInfo"] = flashes[0]
	}
	session.Save()

	var req model.SignInRequest
	if c.Request.Method == http.MethodGet {
		data["Values"] = req
		render(c, http.StatusOK, "sign_in_index.html", data)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		data["Values"] = req
		data["Message"] = "Validation error"
		data["Errors"] = validatorutil.ValidatorError(err)
		render(c, http.StatusBadRequest, "sign_in_index.html", data)
		return
	}

	result, err := h.authService.SignIn(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailNotFound:
			errorsMap["Email"] = "Email not registered"
		case domain.ErrUserPasswordNotMatch:
			errorsMap["Password"] = "Wrong password"
		default:
			errorsMap["Error"] = "Something went wrong"
		}
		data["Values"] = req
		data["Errors"] = errorsMap
		render(c, http.StatusUnauthorized, "sign_in_index.html", data)
		return
	}

	if result == nil {
		errorsMap := make(map[string]string)
		errorsMap["Error"] = "Something went wrong"
		data["Values"] = req
		data["Errors"] = errorsMap
		render(c, http.StatusUnauthorized, "sign_in_index.html", data)
		return
	}

	saveUserSession(session, *result)

	session.AddFlash("Sign in successfully!", "success")
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *authController) SignOut(c *gin.Context) {
	session := sessions.Default(c)

	ClearUserSession(session)

	c.Redirect(http.StatusFound, "/sign-in")
}

func NewAuthController(
	authService service.IAuthService,
) *authController {
	return &authController{
		authService: authService,
	}
}

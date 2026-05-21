package handler

import (
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/model"
	"agussyahrilmubarok.github.io/backend/internal/service"
	"agussyahrilmubarok.github.io/backend/pkg/helper"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.IAuthService
}

// SignUp godoc
// @Summary      Sign up
// @Description  Create new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.SignUpRequest   true  "Sign Up Request"
// @Success      201      {object}  model.SuccessResponse
// @Failure      409      {object}  model.ErrorResponse
// @Failure      422      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/auth/sign-up [post]
func (h *authHandler) SignUp(c *gin.Context) {
	var req model.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Message: "Validation error",
			Errors:  helper.ValidatorError(err),
		})
		return
	}

	result, err := h.authService.SignUp(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailExists:
			errorsMap["email"] = "Email already exists"
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "Conflict error",
				Errors:  errorsMap,
			})
			return
		default:
			errorsMap["error"] = "Failed to sign up"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal server error",
				Errors:  errorsMap,
			})
			return
		}
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Message: "Sign up successfully",
		Data:    result,
	})
}

func (h *authHandler) SignIn(c *gin.Context) {

}

func NewAuthHandler(
	authService service.IAuthService,
) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

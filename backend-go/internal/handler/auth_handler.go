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
// @Param        request  body      model.SignUpRequest  true  "Sign Up Request"
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
		default:
			errorsMap["error"] = "Failed to sign up"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal server error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Message: "Sign up successfully",
		Data:    result,
	})
}

// SignIn godoc
// @Summary      Sign in
// @Description  Authenticate user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.SignInRequest  true  "Sign In Request"
// @Success      200      {object}  model.SuccessResponse
// @Failure      401      {object}  model.ErrorResponse
// @Failure      404      {object}  model.ErrorResponse
// @Failure      422      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/auth/sign-in [post]
func (h *authHandler) SignIn(c *gin.Context) {
	var req model.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Message: "Validation error",
			Errors:  helper.ValidatorError(err),
		})
		return
	}

	result, err := h.authService.SignIn(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailNotFound:
			errorsMap["email"] = "Email not found"
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "Not found error",
				Errors:  errorsMap,
			})
		case domain.ErrUserPasswordNotMatch:
			errorsMap["password"] = "Password not match"
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized error",
				Errors:  errorsMap,
			})
		default:
			errorsMap["error"] = "Failed to sign in"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal server error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Sign in successfully",
		Data:    result,
	})
}

func NewAuthHandler(
	authService service.IAuthService,
) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

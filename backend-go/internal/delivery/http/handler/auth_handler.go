package handler

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/model"
	"agussyahrilmubarok.github.io/backend/pkg/validator"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC usecase.AuthUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

// SignUp godoc
// @Summary      Sign up
// @Description  Register a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.SignUpRequest  true  "Sign up request"
// @Success      201      {object}  model.SuccessResponse
// @Failure      400      {object}  model.ErrorResponse
// @Failure      409      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/auth/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	res, err := h.authUC.SignUp(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailAlreadyInUse) {
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "Email is already in use"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Something went wrong"},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Message: "Signed up successfully",
		Data:    res,
	})
}

// SignIn godoc
// @Summary      Sign in
// @Description  Authenticate user and return token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.SignInRequest  true  "Sign in request"
// @Success      200      {object}  model.SuccessResponse
// @Failure      400      {object}  model.ErrorResponse
// @Failure      401      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	res, err := h.authUC.SignIn(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailNotRegistered) {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"email": "Email is not registered"},
			})
			return
		}
		if errors.Is(err, usecase.ErrPasswordMismatch) {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"password": "Password do not match"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Something went wrong"},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Signed in successfully",
		Data:    res,
	})
}

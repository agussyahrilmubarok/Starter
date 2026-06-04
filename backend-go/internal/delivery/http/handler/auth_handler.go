package handler

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/payload"
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
// @Success      201      {object}  payload.SuccessResponse
// @Failure      400      {object}  payload.ErrorResponse
// @Failure      409      {object}  payload.ErrorResponse
// @Failure      500      {object}  payload.ErrorResponse
// @Router       /v1/auth/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	res, err := h.authUC.SignUp(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, payload.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "The email has already been taken"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Something went wrong"},
		})
		return
	}

	c.JSON(http.StatusCreated, payload.SuccessResponse{
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
// @Success      200      {object}  payload.SuccessResponse
// @Failure      400      {object}  payload.ErrorResponse
// @Failure      401      {object}  payload.ErrorResponse
// @Failure      500      {object}  payload.ErrorResponse
// @Router       /v1/auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	res, err := h.authUC.SignIn(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidEmail) {
			c.JSON(http.StatusUnauthorized, payload.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"email": "The email address is not registered"},
			})
			return
		}
		if errors.Is(err, usecase.ErrInvalidPassword) {
			c.JSON(http.StatusUnauthorized, payload.ErrorResponse{
				Message: "Unauthorized",
				Errors:  map[string]string{"password": "The password is incorrect"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Something went wrong"},
		})
		return
	}

	c.JSON(http.StatusOK, payload.SuccessResponse{
		Message: "Signed in successfully",
		Data:    res,
	})
}

package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/exception"
	"backend/pkg/helper"
	"backend/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("[WARN] failed validate request body: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, response.HttpError{
			Success: false,
			Message: "Validation errors",
			Errors:  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		log.Printf("[ERROR] failed register new user: %v\n", err)
		if ex, ok := err.(*exception.Exception); ok {
			c.JSON(ex.Code, response.HttpError{
				Success: false,
				Message: ex.Message,
				Errors:  ex.Err,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.HttpError{
			Success: false,
			Message: "Internal Server Error",
			Errors:  response.HttpErrMap("Error", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, response.HttpSuccess{
		Success: true,
		Message: "User register successfully",
		Data:    result,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.UserLoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("[WARN] failed validate request body: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, response.HttpError{
			Success: false,
			Message: "Validation errors",
			Errors:  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		log.Printf("[ERROR] failed login user: %v\n", err)
		if ex, ok := err.(*exception.Exception); ok {
			c.JSON(ex.Code, response.HttpError{
				Success: false,
				Message: ex.Message,
				Errors:  ex.Err,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.HttpError{
			Success: false,
			Message: "Internal Server Error",
			Errors:  response.HttpErrMap("Error", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, response.HttpSuccess{
		Success: true,
		Message: "Login Success",
		Data:    result,
	})
}

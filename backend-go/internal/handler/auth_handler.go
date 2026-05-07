package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/usecase"
	"backend/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase usecase.IAuthUseCase
}

func NewAuthHandler(
	authUseCase usecase.IAuthUseCase,
) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Validation Error",
			"errors":  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.authUseCase.Register(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case domain.ErrUserEmailUsed:
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Email is already used",
				"errors": gin.H{
					"email": "Email is already used",
				},
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
				"errors": gin.H{
					"error": "Internal server error",
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register user successfully",
		"data":    result,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Validation errors",
			"errors":  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.authUseCase.Login(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "User not found",
				"errors": gin.H{
					"email": "Email is not registered",
				},
			})
			return
		case domain.ErrUserPasswordInvalid:
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "User not found",
				"errors": gin.H{
					"password": "Password does not match",
				},
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
				"errors": gin.H{
					"error": "Internal server error",
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User login successfully",
		"data":    result,
	})
}

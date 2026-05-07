package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/exception"
	"backend/pkg/helper"
	"backend/pkg/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) FindUsers(c *gin.Context) {
	result, err := h.userService.FindAll(c.Request.Context())
	if err != nil {
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
		Message: "List Data Users",
		Data:    result,
	})
}

func (h *UserHandler) FindUserById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	result, err := h.userService.FindByID(c.Request.Context(), uint(id))
	if err != nil {
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
		Message: "User created successfully",
		Data:    result,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("[ERROR] failed to validate request body: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, response.HttpError{
			Success: false,
			Message: "Validation errors",
			Errors:  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.userService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Printf("[ERROR] failed create new user: %v\n", err)
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
		Message: "User Found",
		Data:    result,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var req model.UserUpdateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Printf("[ERROR] failed to validate request body: %v\n", err)
		c.JSON(http.StatusUnprocessableEntity, response.HttpError{
			Success: false,
			Message: "Validation errors",
			Errors:  helper.ValidateRequest(err),
		})
		return
	}

	result, err := h.userService.UpdateByID(c.Request.Context(), uint(id), &req)
	if err != nil {
		log.Printf("[ERROR] failed update user: %v\n", err)
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
		Message: "User updated successfully",
		Data:    result,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := h.userService.DeleteByID(c.Request.Context(), uint(id)); err != nil {
		log.Printf("[ERROR] failed delete user: %v\n", err)
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
		Message: "User deleted successfully",
	})
}

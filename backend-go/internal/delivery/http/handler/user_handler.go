package handler

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/payload"
	"agussyahrilmubarok.github.io/backend/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUC usecase.UserUseCase
}

func NewUserHandler(userUC usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// GetAll godoc
// @Summary      Get all users
// @Description  Get all users
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  payload.SuccessResponse
// @Failure      500  {object}  payload.ErrorResponse
// @Router       /v1/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userUC.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Failed to get users"},
		})
		return
	}

	c.JSON(http.StatusOK, payload.SuccessResponse{
		Message: "Get all users successfully",
		Data:    users,
	})
}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Get user by ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  payload.SuccessResponse
// @Failure      400  {object}  payload.ErrorResponse
// @Failure      404  {object}  payload.ErrorResponse
// @Failure      500  {object}  payload.ErrorResponse
// @Router       /v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	user, err := h.userUC.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, payload.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Failed to get user"},
		})
		return
	}

	c.JSON(http.StatusOK, payload.SuccessResponse{
		Message: "Get user successfully",
		Data:    user,
	})
}

// Create godoc
// @Summary      Create user
// @Description  Create a new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.CreateUserRequest  true  "Create user request"
// @Success      201      {object}  payload.SuccessResponse
// @Failure      400      {object}  payload.ErrorResponse
// @Failure      409      {object}  payload.ErrorResponse
// @Failure      500      {object}  payload.ErrorResponse
// @Router       /v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	user, err := h.userUC.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, payload.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "Email already exists"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Failed to create user"},
		})
		return
	}

	c.JSON(http.StatusCreated, payload.SuccessResponse{
		Message: "User created successfully",
		Data:    user,
	})
}

// Update godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                 true  "User ID"
// @Param        request  body      dto.UpdateUserRequest  true  "Update user request"
// @Success      200      {object}  payload.SuccessResponse
// @Failure      400      {object}  payload.ErrorResponse
// @Failure      404      {object}  payload.ErrorResponse
// @Failure      409      {object}  payload.ErrorResponse
// @Failure      500      {object}  payload.ErrorResponse
// @Router       /v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	user, err := h.userUC.Update(c.Request.Context(), id, req)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserNotFound):
			c.JSON(http.StatusNotFound, payload.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
			})
		case errors.Is(err, usecase.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, payload.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "Email already exists"},
			})
		default:
			c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
				Message: "Internal server error",
				Errors:  map[string]string{"error": "Failed to update user"},
			})
		}
		return
	}

	c.JSON(http.StatusOK, payload.SuccessResponse{
		Message: "User updated successfully",
		Data:    user,
	})
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  payload.SuccessResponse
// @Failure      400  {object}  payload.ErrorResponse
// @Failure      404  {object}  payload.ErrorResponse
// @Failure      500  {object}  payload.ErrorResponse
// @Router       /v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, payload.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	if err := h.userUC.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, payload.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, payload.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Failed to delete user"},
		})
		return
	}

	c.JSON(http.StatusOK, payload.SuccessResponse{
		Message: "User deleted successfully",
	})
}

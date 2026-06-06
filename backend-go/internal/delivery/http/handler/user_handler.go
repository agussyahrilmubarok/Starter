package handler

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/model"
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
// @Success      200  {object}  model.SuccessResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /v1/users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userUC.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal server error",
			Errors:  map[string]string{"error": "Something went wrong"},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Users retrieved successfully",
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
// @Success      200  {object}  model.SuccessResponse
// @Failure      400  {object}  model.ErrorResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	user, err := h.userUC.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
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
		Message: "User retrieved successfully",
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
// @Success      201      {object}  model.SuccessResponse
// @Failure      400      {object}  model.ErrorResponse
// @Failure      409      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	user, err := h.userUC.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailAlreadyInUse) {
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "Email is already use"},
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
// @Success      200      {object}  model.SuccessResponse
// @Failure      400      {object}  model.ErrorResponse
// @Failure      404      {object}  model.ErrorResponse
// @Failure      409      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  validator.ParseError(err),
		})
		return
	}

	user, err := h.userUC.Update(c.Request.Context(), id, req)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserNotFound):
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
			})
		case errors.Is(err, usecase.ErrEmailAlreadyInUse):
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "Conflict",
				Errors:  map[string]string{"email": "Email is already use"},
			})
		default:
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal server error",
				Errors:  map[string]string{"error": "Something went wrong"},
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
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
// @Success      200  {object}  model.SuccessResponse
// @Failure      400  {object}  model.ErrorResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Bad request",
			Errors:  map[string]string{"error": "Invalid user id"},
		})
		return
	}

	if err := h.userUC.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "Not found",
				Errors:  map[string]string{"error": "User not found"},
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
		Message: "User deleted successfully",
	})
}

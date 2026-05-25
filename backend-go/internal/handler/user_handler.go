package handler

import (
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/model"
	"agussyahrilmubarok.github.io/backend/internal/service"
	"agussyahrilmubarok.github.io/backend/pkg/validatorutil"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.IUserService
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
func (h *userHandler) GetAll(c *gin.Context) {
	result, err := h.userService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "Internal Server Error",
			Errors: map[string]string{
				"error": "Failed to get users",
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Get All Users Successfully",
		Data:    result,
	})
}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Get user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  model.SuccessResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /v1/users/{id} [get]
func (h *userHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserNotFound:
			errorsMap["id"] = "User not found"
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "User Not Found",
				Errors:  errorsMap,
			})
		default:
			errorsMap["error"] = "Failed to get user"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal Server Error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Get User Successfully",
		Data:    result,
	})
}

// Create godoc
// @Summary      Create user
// @Description  Create new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateUserRequest  true  "Create User Request"
// @Security     BearerAuth
// @Success      201      {object}  model.SuccessResponse
// @Failure      409      {object}  model.ErrorResponse
// @Failure      422      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/users [post]
func (h *userHandler) Create(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Message: "Validation Failed",
			Errors:  validatorutil.ValidatorError(err),
		})
		return
	}

	result, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailExists:
			errorsMap["email"] = "Email already exists"
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "User Email Already Exists",
				Errors:  errorsMap,
			})
		default:
			errorsMap["error"] = "Failed to create user"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal Server Error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Message: "Create User Successfully",
		Data:    result,
	})
}

// UpdateByID godoc
// @Summary      Update user by ID
// @Description  Update user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id       path      string                   true  "User ID"
// @Param        request  body      model.UpdateUserRequest  true  "Update User Request"
// @Security     BearerAuth
// @Success      200      {object}  model.SuccessResponse
// @Failure      404      {object}  model.ErrorResponse
// @Failure      422      {object}  model.ErrorResponse
// @Failure      500      {object}  model.ErrorResponse
// @Router       /v1/users/{id} [put]
func (h *userHandler) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.ErrorResponse{
			Message: "Validation Failed",
			Errors:  validatorutil.ValidatorError(err),
		})
		return
	}

	result, err := h.userService.UpdateByID(c.Request.Context(), id, req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserNotFound:
			errorsMap["id"] = "User not found"
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "User Not Found",
				Errors:  errorsMap,
			})
		case domain.ErrUserEmailExists:
			errorsMap["email"] = "Email already exists"
			c.JSON(http.StatusConflict, model.ErrorResponse{
				Message: "User Email Already Exists",
				Errors:  errorsMap,
			})
		default:
			errorsMap["error"] = "Failed to update user"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal Server Error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Update User Successfully",
		Data:    result,
	})
}

// DeleteByID godoc
// @Summary      Delete user by ID
// @Description  Delete user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Security     BearerAuth
// @Success      200  {object}  model.SuccessResponse
// @Failure      404  {object}  model.ErrorResponse
// @Failure      500  {object}  model.ErrorResponse
// @Router       /v1/users/{id} [delete]
func (h *userHandler) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	if err := h.userService.DeleteByID(c.Request.Context(), id); err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserNotFound:
			errorsMap["id"] = "User not found"
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Message: "User Not Found",
				Errors:  errorsMap,
			})
		default:
			errorsMap["error"] = "Failed to delete user"
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				Message: "Internal Server Error",
				Errors:  errorsMap,
			})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Delete User Successfully",
	})
}

func NewUserHandler(
	userService service.IUserService,
) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

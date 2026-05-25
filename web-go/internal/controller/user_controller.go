// internal/controller/user_controller.go
package controller

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/model"
	"agussyahrilmubarok.github.io/web/internal/service"
	"agussyahrilmubarok.github.io/web/pkg/helper"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.IUserService
}

func (h *userController) Index(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	session := sessions.Default(c)
	data := gin.H{
		"Title":       "Users",
		"UserProfile": userProfile,
	}

	if flashes := session.Flashes("success"); len(flashes) > 0 {
		data["MsgInfo"] = flashes[0]
	}
	session.Save()

	users, err := h.userService.GetAll(c.Request.Context())
	if err != nil {
		data["MsgError"] = "Failed to load users"
		render(c, http.StatusInternalServerError, "users_index.html", data)
		return
	}

	data["Values"] = users
	render(c, http.StatusOK, "users_index.html", data)
}

func (h *userController) Create(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	data := gin.H{
		"Title":       "Create User",
		"UserProfile": userProfile,
	}

	var req model.CreateUserRequest
	data["Values"] = req

	render(c, http.StatusOK, "users_create.html", data)
}

func (h *userController) Store(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	data := gin.H{
		"Title":       "Create User",
		"UserProfile": userProfile,
	}

	var req model.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		data["Values"] = req
		data["Errors"] = helper.ValidatorError(err)
		render(c, http.StatusBadRequest, "users_create.html", data)
		return
	}

	_, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserEmailExists:
			errorsMap["Email"] = "Email already exists"
		default:
			errorsMap["Error"] = "Something went wrong"
		}
		data["Values"] = req
		data["Errors"] = errorsMap
		render(c, http.StatusBadRequest, "users_create.html", data)
		return
	}

	session := sessions.Default(c)
	session.AddFlash("User created successfully!", "success")
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

// Edit — GET /dashboard/users/:id/edit
func (h *userController) Edit(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	id := c.Param("id")
	user, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil || user == nil {
		session := sessions.Default(c)
		session.AddFlash("User not found", "error")
		session.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	data := gin.H{
		"Title":       "Edit User",
		"UserProfile": userProfile,
		"Data":        user,
	}
	render(c, http.StatusOK, "users_edit.html", data)
}

// Update — POST /dashboard/users/:id/edit
func (h *userController) Update(c *gin.Context) {
	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}

	id := c.Param("id")

	var req model.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		user, _ := h.userService.GetByID(c.Request.Context(), id)
		data := gin.H{
			"Title":       "Edit User",
			"UserProfile": userProfile,
			"Data":        user,
			"Errors":      helper.ValidatorError(err),
		}
		render(c, http.StatusBadRequest, "users_edit.html", data)
		return
	}

	_, err := h.userService.UpdateByID(c.Request.Context(), id, req)
	if err != nil {
		errorsMap := make(map[string]string)
		switch err {
		case domain.ErrUserNotFound:
			errorsMap["Error"] = "User not found"
		case domain.ErrUserEmailExists:
			errorsMap["Email"] = "Email already exists"
		default:
			errorsMap["Error"] = "Something went wrong"
		}
		user, _ := h.userService.GetByID(c.Request.Context(), id)
		data := gin.H{
			"Title":       "Edit User",
			"UserProfile": userProfile,
			"Data":        user,
			"Errors":      errorsMap,
		}
		render(c, http.StatusBadRequest, "users_edit.html", data)
		return
	}

	session := sessions.Default(c)
	session.AddFlash("User updated successfully!", "success")
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

// Delete — POST /dashboard/users/:id/delete
func (h *userController) Delete(c *gin.Context) {
	id := c.Param("id")

	session := sessions.Default(c)

	if err := h.userService.DeleteByID(c.Request.Context(), id); err != nil {
		session.AddFlash("Failed to delete user", "error")
	} else {
		session.AddFlash("User deleted successfully!", "success")
	}
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

func NewUserController(userService service.IUserService) *userController {
	return &userController{userService: userService}
}

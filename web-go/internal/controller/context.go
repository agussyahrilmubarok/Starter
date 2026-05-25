package controller

import (
	"errors"
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotInContext = errors.New("user not found in context")
)

func getUserFromContext(c *gin.Context) (model.UserResponse, error) {
	val, exists := c.Get("user")
	if !exists {
		return model.UserResponse{}, ErrUserNotInContext
	}

	user, ok := val.(model.UserResponse)
	if !ok {
		session := sessions.Default(c)
		ClearUserSession(session)
		return model.UserResponse{}, ErrUserNotInContext
	}

	return user, nil
}

func mustGetUserFromContext(c *gin.Context) (model.UserResponse, bool) {
	user, err := getUserFromContext(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/sign-in")
		c.Abort()
		return model.UserResponse{}, false
	}

	return user, true
}

package controller

import (
	"net/http"
	"strings"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/payload"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/pkg/crypto"
	"agussyahrilmubarok.github.io/web/pkg/i18n"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"agussyahrilmubarok.github.io/web/pkg/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *AppController) UserListPage(c *gin.Context) {
	data := gin.H{
		"Title": "Users",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	log := logger.FromCtx(c.Request.Context())

	s := sessions.Default(c)
	if flashes := s.Flashes("success"); len(flashes) > 0 {
		data["MsgSuccess"] = flashes[0]
	}
	if flashes := s.Flashes("error"); len(flashes) > 0 {
		data["MsgError"] = flashes[0]
	}
	s.Save()

	users, err := h.userRepository.FindAll(c.Request.Context())
	if err != nil {
		log.Error("failed to get all users", zap.Error(err))
		data["MsgError"] = i18n.T(getLang(c), "error.general")
		render(c, http.StatusInternalServerError, "users_index.html", data)
		return
	}

	data["Values"] = payload.ToUserListResponse(users)
	render(c, http.StatusOK, "users_index.html", data)
}

func (h *AppController) UserCreatePage(c *gin.Context) {
	data := gin.H{
		"Title": "Create User",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile
	data["Values"] = payload.CreateUserRequest{}

	render(c, http.StatusOK, "users_create.html", data)
}

func (h *AppController) UserCreate(c *gin.Context) {
	data := gin.H{
		"Title": "Create User",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	log := logger.FromCtx(c.Request.Context())
	lang := getLang(c)

	var req payload.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Warn("failed to validate request", zap.Error(err))
		data["Values"] = req
		data["Errors"] = validator.ParseError(err)
		render(c, http.StatusBadRequest, "users_create.html", data)
		return
	}

	exists, err := h.userRepository.ExistsByEmail(c.Request.Context(), strings.ToLower(req.Email))
	if err != nil {
		log.Error("failed to check email existence", zap.Error(err))
		data["Values"] = req
		data["MsgError"] = i18n.T(lang, "auth.general.error")
		render(c, http.StatusInternalServerError, "users_create.html", data)
		return
	}
	if exists {
		data["Values"] = req
		data["Errors"] = map[string]string{"Email": i18n.T(lang, "user.email.alreadyTaken")}
		render(c, http.StatusBadRequest, "users_create.html", data)
		return
	}

	hashed, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		data["Values"] = req
		data["MsgError"] = i18n.T(lang, "auth.general.error")
		render(c, http.StatusInternalServerError, "users_create.html", data)
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    strings.ToLower(req.Email),
		Password: hashed,
	}

	if err := h.userRepository.Create(c.Request.Context(), &user); err != nil {
		log.Error("failed to create user", zap.Error(err))
		data["Values"] = req
		data["MsgError"] = i18n.T(lang, "auth.general.error")
		render(c, http.StatusInternalServerError, "users_create.html", data)
		return
	}

	s := sessions.Default(c)
	s.AddFlash(i18n.T(lang, "user.create.success"), "success")
	s.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

func (h *AppController) UserEditPage(c *gin.Context) {
	data := gin.H{
		"Title": "Edit User",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	lang := getLang(c)
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	user, err := h.userRepository.FindByID(c.Request.Context(), parsedID)
	if err != nil || user == nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}
	data["Values"] = payload.ToUserResponse(user)

	render(c, http.StatusOK, "users_edit.html", data)
}

func (h *AppController) UserEdit(c *gin.Context) {
	data := gin.H{
		"Title": "Edit User",
	}

	userProfile, ok := mustGetUserFromContext(c)
	if !ok {
		return
	}
	data["UserProfile"] = userProfile

	log := logger.FromCtx(c.Request.Context())
	lang := getLang(c)

	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	user, err := h.userRepository.FindByID(c.Request.Context(), parsedID)
	if err != nil || user == nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	var req payload.UpdateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Warn("failed to validate request", zap.Error(err))
		data["Values"] = payload.ToUserResponse(user)
		data["Errors"] = validator.ParseError(err)
		render(c, http.StatusBadRequest, "users_edit.html", data)
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" && strings.ToLower(req.Email) != user.Email {
		exists, err := h.userRepository.ExistsByEmail(c.Request.Context(), strings.ToLower(req.Email))
		if err != nil {
			log.Error("failed to check email existence", zap.Error(err))
			data["Values"] = payload.ToUserResponse(user)
			data["MsgError"] = i18n.T(lang, "auth.general.error")
			render(c, http.StatusInternalServerError, "users_edit.html", data)
			return
		}
		if exists {
			data["Values"] = payload.ToUserResponse(user)
			data["Errors"] = map[string]string{"Email": i18n.T(lang, "user.email.alreadyTaken")}
			render(c, http.StatusBadRequest, "users_edit.html", data)
			return
		}
		user.Email = strings.ToLower(req.Email)
	}

	if req.Password != "" {
		hashed, err := crypto.HashPassword(req.Password)
		if err != nil {
			log.Error("failed to hash password", zap.Error(err))
			data["Values"] = payload.ToUserResponse(user)
			data["MsgError"] = i18n.T(lang, "auth.general.error")
			render(c, http.StatusInternalServerError, "users_edit.html", data)
			return
		}
		user.Password = hashed
	}

	if err := h.userRepository.Update(c.Request.Context(), user); err != nil {
		log.Error("failed to update user", zap.Error(err))
		data["Values"] = payload.ToUserResponse(user)
		data["MsgError"] = i18n.T(lang, "auth.general.error")
		render(c, http.StatusInternalServerError, "users_edit.html", data)
		return
	}

	s := sessions.Default(c)
	s.AddFlash(i18n.T(lang, "user.update.success"), "success")
	s.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

func (h *AppController) UserDelete(c *gin.Context) {
	lang := getLang(c)
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	log := logger.FromCtx(c.Request.Context())

	user, err := h.userRepository.FindByID(c.Request.Context(), parsedID)
	if err != nil || user == nil {
		s := sessions.Default(c)
		s.AddFlash(i18n.T(lang, "user.notFound"), "error")
		s.Save()
		c.Redirect(http.StatusFound, "/dashboard/users")
		return
	}

	s := sessions.Default(c)
	if err := h.userRepository.Delete(c.Request.Context(), user.ID); err != nil {
		log.Error("failed to delete user", zap.Error(err))
		s.AddFlash(i18n.T(lang, "error.general"), "error")
	} else {
		log.Info("user deleted", zap.String("id", user.ID.String()))
		s.AddFlash(i18n.T(lang, "user.delete.success"), "success")
	}
	s.Save()

	c.Redirect(http.StatusFound, "/dashboard/users")
}

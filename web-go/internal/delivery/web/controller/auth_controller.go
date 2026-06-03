package controller

import (
	"net/http"
	"strings"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/payload"
	"agussyahrilmubarok.github.io/web/internal/delivery/web/session"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/pkg/crypto"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"agussyahrilmubarok.github.io/web/pkg/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *AppController) SignUp(c *gin.Context) {
	data := gin.H{
		"Title": "Sign Up",
	}

	log := logger.FromCtx(c.Request.Context())

	s := sessions.Default(c)
	if flashes := s.Flashes("success"); len(flashes) > 0 {
		data["MsgInfo"] = flashes[0]
	}
	s.Save()

	var req payload.SignUpRequest

	if c.Request.Method == http.MethodGet {
		data["Values"] = req
		render(c, http.StatusOK, "sign_up_index.html", data)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		log.Warn("failed to validate request", zap.Error(err))
		data["Values"] = req
		data["Message"] = "Validation error"
		data["Errors"] = validator.ParseError(err)
		render(c, http.StatusBadRequest, "sign_up_index.html", data)
		return
	}

	exists, err := h.userRepository.ExistsByEmail(c.Request.Context(), strings.ToLower(req.Email))
	if err != nil {
		log.Error("failed to check email existence", zap.Error(err))
		data["Values"] = req
		data["Message"] = "Internal server error"
		data["MsgError"] = "Something went wrong"
		render(c, http.StatusInternalServerError, "sign_up_index.html", data)
		return
	}
	if exists {
		log.Warn("email already exists")
		data["Values"] = req
		data["Message"] = "Bad request"
		data["Errors"] = map[string]string{"Email": "Email already exists"}
		render(c, http.StatusBadRequest, "sign_up_index.html", data)
		return
	}

	hashed, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		data["Values"] = req
		data["Message"] = "Internal server error"
		data["MsgError"] = "Something went wrong"
		render(c, http.StatusInternalServerError, "sign_up_index.html", data)
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
		data["Message"] = "Internal server error"
		data["MsgError"] = "Something went wrong"
		render(c, http.StatusInternalServerError, "sign_up_index.html", data)
		return
	}

	s.AddFlash("Sign up successfully!", "success")
	s.Save()

	c.Redirect(http.StatusFound, "/sign-in")
}

func (h *AppController) SignIn(c *gin.Context) {
	data := gin.H{
		"Title": "Sign In",
	}

	log := logger.FromCtx(c.Request.Context())

	s := sessions.Default(c)
	if flashes := s.Flashes("success"); len(flashes) > 0 {
		data["MsgInfo"] = flashes[0]
	}
	s.Save()

	var req payload.SignInRequest

	if c.Request.Method == http.MethodGet {
		data["Values"] = req
		render(c, http.StatusOK, "sign_in_index.html", data)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		log.Warn("failed to validate request", zap.Error(err))
		data["Values"] = req
		data["Message"] = "Validation error"
		data["Errors"] = validator.ParseError(err)
		render(c, http.StatusBadRequest, "sign_in_index.html", data)
		return
	}

	user, err := h.userRepository.FindByEmail(c.Request.Context(), strings.ToLower(req.Email))
	if err != nil {
		log.Error("failed to find user", zap.Error(err))
		data["Values"] = req
		data["Message"] = "Internal server error"
		data["MsgError"] = "Something went wrong"
		render(c, http.StatusInternalServerError, "sign_in_index.html", data)
		return
	}
	if user == nil {
		log.Warn("invalid email")
		data["Values"] = req
		data["Message"] = "Bad request"
		data["Errors"] = map[string]string{"Email": "Invalid email"}
		render(c, http.StatusBadRequest, "sign_in_index.html", data)
		return
	}

	if !crypto.CheckPassword(user.Password, req.Password) {
		log.Warn("invalid password")
		data["Values"] = req
		data["Message"] = "Bad request"
		data["Errors"] = map[string]string{"Password": "Invalid password"}
		render(c, http.StatusBadRequest, "sign_in_index.html", data)
		return
	}

	session.SaveUser(s, payload.ToUserResponse(user))

	s.AddFlash("Sign in successfully!", "success")
	s.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *AppController) SignOut(c *gin.Context) {
	s := sessions.Default(c)

	session.DeleteUser(s)

	c.Redirect(http.StatusFound, "/sign-in")
}

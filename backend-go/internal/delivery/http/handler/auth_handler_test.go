package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ucmocks "agussyahrilmubarok.github.io/backend/internal/application/usecase/mocks"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupAuthRouter(h *handler.AuthHandler) *gin.Engine {
	r := gin.New()
	r.POST("/v1/auth/sign-up", h.SignUp)
	r.POST("/v1/auth/sign-in", h.SignIn)
	return r
}

func TestAuthHandler_SignUp_Success(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignUpRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	authResp := &dto.AuthResponse{
		Token: "jwt-token",
		User: dto.UserResponse{
			Name:  "John Doe",
			Email: "john@example.com",
		},
	}

	authUC.On("SignUp", mock.Anything, reqBody).Return(authResp, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-up", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Signed up successfully", resp["message"])
	assert.NotNil(t, resp["data"])
}

func TestAuthHandler_SignUp_InvalidBody(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	body := []byte(`{"name":"Jo"}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-up", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Bad request", resp["message"])
}

func TestAuthHandler_SignUp_EmailAlreadyExist(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignUpRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	authUC.On("SignUp", mock.Anything, reqBody).Return(nil, usecase.ErrEmailAlreadyExists)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-up", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Conflict", resp["message"])
}

func TestAuthHandler_SignUp_InternalError(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignUpRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	authUC.On("SignUp", mock.Anything, reqBody).Return(nil, errors.New("db error"))

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-up", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Internal server error", resp["message"])
}

func TestAuthHandler_SignIn_Success(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignInRequest{
		Email:    "john@example.com",
		Password: "password123",
	}
	authResp := &dto.AuthResponse{
		Token: "jwt-token",
		User: dto.UserResponse{
			Name:  "John Doe",
			Email: "john@example.com",
		},
	}

	authUC.On("SignIn", mock.Anything, reqBody).Return(authResp, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Signed in successfully", resp["message"])
}

func TestAuthHandler_SignIn_InvalidBody(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	body := []byte(`{}`) // missing email & password
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthHandler_SignIn_InvalidEmail(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignInRequest{
		Email:    "notfound@example.com",
		Password: "password123",
	}

	authUC.On("SignIn", mock.Anything, reqBody).Return(nil, usecase.ErrInvalidEmail)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Unauthorized", resp["message"])
	errs := resp["errors"].(map[string]any)
	assert.Equal(t, "The email address is not registered", errs["email"])
}

func TestAuthHandler_SignIn_InvalidPassword(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignInRequest{
		Email:    "john@example.com",
		Password: "wrongpass",
	}

	authUC.On("SignIn", mock.Anything, reqBody).Return(nil, usecase.ErrInvalidPassword)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	errs := resp["errors"].(map[string]any)
	assert.Equal(t, "The password is incorrect", errs["password"])
}

func TestAuthHandler_SignIn_InternalError(t *testing.T) {
	authUC := ucmocks.NewAuthUseCase(t)
	h := handler.NewAuthHandler(authUC)
	r := setupAuthRouter(h)

	reqBody := dto.SignInRequest{
		Email:    "john@example.com",
		Password: "password123",
	}

	authUC.On("SignIn", mock.Anything, reqBody).Return(nil, errors.New("db error"))

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

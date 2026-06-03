package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	ucmocks "agussyahrilmubarok.github.io/backend/internal/application/usecase/mocks"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupUserRouter(h *handler.UserHandler) *gin.Engine {
	r := gin.New()
	r.GET("/v1/users", h.GetAll)
	r.GET("/v1/users/:id", h.GetByID)
	r.POST("/v1/users", h.Create)
	r.PUT("/v1/users/:id", h.Update)
	r.DELETE("/v1/users/:id", h.Delete)
	return r
}

func newSampleUser() dto.UserResponse {
	return dto.UserResponse{
		ID:        uuid.New(),
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestUserHandler_GetAll_Success(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	users := []dto.UserResponse{newSampleUser(), newSampleUser()}
	userUC.On("GetAll", mock.Anything).Return(users, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Get all users successfully", resp["message"])
	data := resp["data"].([]any)
	assert.Len(t, data, 2)
}

func TestUserHandler_GetAll_InternalError(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	userUC.On("GetAll", mock.Anything).Return(nil, errors.New("db error"))

	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Internal server error", resp["message"])
}

func TestUserHandler_GetByID_Success(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	user := newSampleUser()
	userUC.On("GetByID", mock.Anything, user.ID).Return(&user, nil)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", user.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Get user successfully", resp["message"])
}

func TestUserHandler_GetByID_InvalidUUID(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	req := httptest.NewRequest(http.MethodGet, "/v1/users/not-a-uuid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	errs := resp["errors"].(map[string]any)
	assert.Equal(t, "Invalid user id", errs["error"])
}

func TestUserHandler_GetByID_NotFound(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	userUC.On("GetByID", mock.Anything, id).Return(nil, usecase.ErrUserNotFound)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Not found", resp["message"])
}

func TestUserHandler_GetByID_InternalError(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	userUC.On("GetByID", mock.Anything, id).Return(nil, errors.New("db error"))

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_Create_Success(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	reqBody := dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	user := newSampleUser()
	userUC.On("Create", mock.Anything, reqBody).Return(&user, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "User created successfully", resp["message"])
}

func TestUserHandler_Create_InvalidBody(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	body := []byte(`{"name":"J"}`) // name terlalu pendek, email & password missing
	req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_Create_EmailAlreadyExist(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	reqBody := dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	userUC.On("Create", mock.Anything, reqBody).Return(nil, usecase.ErrEmailAlreadyExists)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Conflict", resp["message"])
}

func TestUserHandler_Create_InternalError(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	reqBody := dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	userUC.On("Create", mock.Anything, reqBody).Return(nil, errors.New("db error"))

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_Update_Success(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	reqBody := dto.UpdateUserRequest{Name: "Jane Doe"}
	user := newSampleUser()
	user.Name = "Jane Doe"

	userUC.On("Update", mock.Anything, id, reqBody).Return(&user, nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/users/%s", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "User updated successfully", resp["message"])
}

func TestUserHandler_Update_InvalidUUID(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	body := []byte(`{"name":"Jane"}`)
	req := httptest.NewRequest(http.MethodPut, "/v1/users/not-a-uuid", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_Update_NotFound(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	reqBody := dto.UpdateUserRequest{Name: "Jane Doe"}

	userUC.On("Update", mock.Anything, id, reqBody).Return(nil, usecase.ErrUserNotFound)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/users/%s", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUserHandler_Update_EmailAlreadyExist(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	reqBody := dto.UpdateUserRequest{Email: "taken@example.com"}

	userUC.On("Update", mock.Anything, id, reqBody).Return(nil, usecase.ErrEmailAlreadyExists)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/users/%s", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestUserHandler_Update_InternalError(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	reqBody := dto.UpdateUserRequest{Name: "Jane Doe"}

	userUC.On("Update", mock.Anything, id, reqBody).Return(nil, errors.New("db error"))

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/v1/users/%s", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_Delete_Success(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	userUC.On("Delete", mock.Anything, id).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/users/%s", id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "User deleted successfully", resp["message"])
}

func TestUserHandler_Delete_InvalidUUID(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	req := httptest.NewRequest(http.MethodDelete, "/v1/users/not-a-uuid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_Delete_NotFound(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	userUC.On("Delete", mock.Anything, id).Return(usecase.ErrUserNotFound)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/users/%s", id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUserHandler_Delete_InternalError(t *testing.T) {
	userUC := ucmocks.NewUserUseCase(t)
	h := handler.NewUserHandler(userUC)
	r := setupUserRouter(h)

	id := uuid.New()
	userUC.On("Delete", mock.Anything, id).Return(errors.New("db error"))

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/users/%s", id), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

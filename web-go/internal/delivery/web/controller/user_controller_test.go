package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/controller"
	"agussyahrilmubarok.github.io/web/internal/delivery/web/payload"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/domain/mocks"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupUserRouter(ac *controller.AppController) *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("test-session", cookie.NewStore([]byte("secret"))))
	r.HTMLRender = &dummyRenderer{}

	injectUser := func(c *gin.Context) {
		c.Set("user", payload.UserResponse{
			ID:    uuid.New().String(),
			Name:  "Admin",
			Email: "admin@example.com",
		})
		c.Next()
	}

	r.GET("/dashboard/users", injectUser, ac.UserListPage)
	r.GET("/dashboard/users/create", injectUser, ac.UserCreatePage)
	r.POST("/dashboard/users/create", injectUser, ac.UserCreate)
	r.GET("/dashboard/users/:id/edit", injectUser, ac.UserEditPage)
	r.POST("/dashboard/users/:id/edit", injectUser, ac.UserEdit)
	r.POST("/dashboard/users/:id/delete", injectUser, ac.UserDelete)

	return r
}

func TestUserController_UserListPage_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	users := []domain.User{*newSampleDomainUser(), *newSampleDomainUser()}
	repo.On("FindAll", mock.Anything).Return(users, nil)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_UserListPage_NoUser_Redirects(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)

	r := gin.New()
	r.Use(sessions.Sessions("test-session", cookie.NewStore([]byte("secret"))))
	r.HTMLRender = &dummyRenderer{}
	r.GET("/dashboard/users", ac.UserListPage)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/sign-in", w.Header().Get("Location"))
}

func TestUserController_UserAddPage_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users/create", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_UserAdd_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	repo.On("ExistsByEmail", mock.Anything, "new@example.com").Return(false, nil)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	body := formBody(map[string]string{
		"name":     "New User",
		"email":    "new@example.com",
		"password": "password123",
	})
	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/create", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEditPage_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	user := newSampleDomainUser()
	repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users/"+user.ID.String()+"/edit", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_UserEditPage_InvalidUUID(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users/not-a-uuid/edit", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEditPage_UserNotFound(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	id := uuid.New()
	repo.On("FindByID", mock.Anything, id).Return(nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users/"+id.String()+"/edit", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEditPage_RepoError(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	id := uuid.New()
	repo.On("FindByID", mock.Anything, id).Return(nil, errors.New("db error"))

	req := httptest.NewRequest(http.MethodGet, "/dashboard/users/"+id.String()+"/edit", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEdit_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	user := newSampleDomainUser()
	repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
	repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	body := formBody(map[string]string{"name": "Updated Name"})
	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/"+user.ID.String()+"/edit", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEdit_InvalidUUID(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	body := formBody(map[string]string{"name": "Test"})
	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/bad-uuid/edit", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserEdit_UserNotFound(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	id := uuid.New()
	repo.On("FindByID", mock.Anything, id).Return(nil, nil)

	body := formBody(map[string]string{"name": "Test"})
	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/"+id.String()+"/edit", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserDelete_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	user := newSampleDomainUser()
	repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
	repo.On("Delete", mock.Anything, user.ID).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/"+user.ID.String()+"/delete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserDelete_InvalidUUID(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/bad-uuid/delete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserDelete_UserNotFound(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	id := uuid.New()
	repo.On("FindByID", mock.Anything, id).Return(nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/"+id.String()+"/delete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

func TestUserController_UserDelete_RepoError(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupUserRouter(ac)

	user := newSampleDomainUser()
	repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
	repo.On("Delete", mock.Anything, user.ID).Return(errors.New("db error"))

	req := httptest.NewRequest(http.MethodPost, "/dashboard/users/"+user.ID.String()+"/delete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard/users", w.Header().Get("Location"))
}

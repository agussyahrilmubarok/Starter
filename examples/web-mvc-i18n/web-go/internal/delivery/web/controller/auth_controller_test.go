package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/controller"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/domain/mocks"
	"agussyahrilmubarok.github.io/web/pkg/crypto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(ac *controller.AppController) *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("test-session", cookie.NewStore([]byte("secret"))))

	r.HTMLRender = &dummyRenderer{}

	r.GET("/sign-up", ac.SignUp)
	r.POST("/sign-up", ac.SignUp)
	r.GET("/sign-in", ac.SignIn)
	r.POST("/sign-in", ac.SignIn)
	r.POST("/sign-out", ac.SignOut)

	return r
}

type dummyRenderer struct{}

func (d *dummyRenderer) Instance(name string, data any) render.Render {
	return &dummyInstance{code: http.StatusOK}
}

type dummyInstance struct{ code int }

func (d *dummyInstance) Render(w http.ResponseWriter) error {
	w.WriteHeader(d.code)
	return nil
}
func (d *dummyInstance) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func formBody(values map[string]string) *strings.Reader {
	form := url.Values{}
	for k, v := range values {
		form.Set(k, v)
	}
	return strings.NewReader(form.Encode())
}

func newSampleDomainUser() *domain.User {
	return &domain.User{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "john@example.com",
	}
}

func TestAuthController_SignUp_GET(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	req := httptest.NewRequest(http.MethodGet, "/sign-up", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthController_SignUp_POST_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	repo.On("ExistsByEmail", mock.Anything, "john@example.com").Return(false, nil)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	body := formBody(map[string]string{
		"name":     "John Doe",
		"email":    "john@example.com",
		"password": "password123",
	})
	req := httptest.NewRequest(http.MethodPost, "/sign-up", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/sign-in", w.Header().Get("Location"))
}

func TestAuthController_SignUp_POST_EmailLowercased(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	repo.On("ExistsByEmail", mock.Anything, "john@example.com").Return(false, nil)
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	body := formBody(map[string]string{
		"name":     "John Doe",
		"email":    "JOHN@EXAMPLE.COM",
		"password": "password123",
	})
	req := httptest.NewRequest(http.MethodPost, "/sign-up", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/sign-in", w.Header().Get("Location"))
}

func TestAuthController_SignIn_GET(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	req := httptest.NewRequest(http.MethodGet, "/sign-in", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthController_SignIn_POST_Success(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	user := newSampleDomainUser()
	hashed, _ := crypto.HashPassword("password123")
	user.Password = hashed

	repo.On("FindByEmail", mock.Anything, "john@example.com").Return(user, nil)

	body := formBody(map[string]string{
		"email":    "john@example.com",
		"password": "password123",
	})
	req := httptest.NewRequest(http.MethodPost, "/sign-in", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/dashboard", w.Header().Get("Location"))
}

func TestAuthController_SignOut(t *testing.T) {
	repo := mocks.NewUserRepository(t)
	ac := controller.NewAppController(repo)
	r := setupRouter(ac)

	req := httptest.NewRequest(http.MethodPost, "/sign-out", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/sign-in", w.Header().Get("Location"))
}

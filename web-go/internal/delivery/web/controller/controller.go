package controller

import (
	"net/http"
	"os"
	"path/filepath"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/payload"
	"agussyahrilmubarok.github.io/web/internal/delivery/web/session"
	"agussyahrilmubarok.github.io/web/internal/domain"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AppController struct {
	userRepository domain.UserRepository
}

func NewAppController(
	userRepository domain.UserRepository,
) *AppController {
	return &AppController{
		userRepository: userRepository,
	}
}

func render(c *gin.Context, code int, tmpl string, data gin.H) {
	c.HTML(code, tmpl, data)
}

func mustGetUserFromContext(c *gin.Context) (payload.UserResponse, bool) {
	val, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/sign-in")
		c.Abort()
		return payload.UserResponse{}, false
	}

	user, ok := val.(payload.UserResponse)
	if !ok {
		s := sessions.Default(c)
		session.DeleteUser(s)
		c.Redirect(http.StatusFound, "/sign-in")
		c.Abort()
		return payload.UserResponse{}, false
	}

	return user, true
}

func (h *AppController) LoadTemplate(templateDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	commons, err := filepath.Glob(templateDir + "/common/*.html")
	if err != nil {
		panic(err.Error())
	}

	homePages, err := filepath.Glob(templateDir + "/home/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, page := range homePages {
		if fileInfo, err := os.Stat(page); err == nil && !fileInfo.IsDir() {
			files := append([]string{filepath.Join(templateDir, "layouts", "default_layout.html")}, page)
			files = append(files, commons...)
			templateName := filepath.Base(page)
			renderer.AddFromFiles(templateName, files...)
		}
	}

	authPages, err := filepath.Glob(templateDir + "/auth/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, page := range authPages {
		if fileInfo, err := os.Stat(page); err == nil && !fileInfo.IsDir() {
			files := append([]string{filepath.Join(templateDir, "layouts", "default_layout.html")}, page)
			files = append(files, commons...)
			templateName := filepath.Base(page)
			renderer.AddFromFiles(templateName, files...)
		}
	}

	dashboardPages, err := filepath.Glob(templateDir + "/dashboard/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, page := range dashboardPages {
		if fileInfo, err := os.Stat(page); err == nil && !fileInfo.IsDir() {
			files := append([]string{filepath.Join(templateDir, "layouts", "dashboard_layout.html")}, page)
			files = append(files, commons...)
			templateName := filepath.Base(page)
			renderer.AddFromFiles(templateName, files...)
		}
	}

	return renderer
}

package controller

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

var templateFuncMap = template.FuncMap{
	"hasPrefix": strings.HasPrefix,
	"not": func(v interface{}) bool {
		if v == nil {
			return true
		}
		return false
	},
}

func render(c *gin.Context, code int, tmpl string, data gin.H) {
	if lang, exists := c.Get("Lang"); exists {
		data["Lang"] = lang
	}
	if t, exists := c.Get("T"); exists {
		data["T"] = t
	}
	data["RequestURI"] = c.Request.URL.Path
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
			files := []string{filepath.Join(templateDir, "layouts", "default_layout.html"), page}
			files = append(files, commons...)
			renderer.AddFromFilesFuncs(filepath.Base(page), templateFuncMap, files...)
		}
	}

	authPages, err := filepath.Glob(templateDir + "/auth/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, page := range authPages {
		if fileInfo, err := os.Stat(page); err == nil && !fileInfo.IsDir() {
			files := []string{filepath.Join(templateDir, "layouts", "default_layout.html"), page}
			files = append(files, commons...)
			renderer.AddFromFilesFuncs(filepath.Base(page), templateFuncMap, files...)
		}
	}

	dashboardPages, err := filepath.Glob(templateDir + "/dashboard/**/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, page := range dashboardPages {
		if fileInfo, err := os.Stat(page); err == nil && !fileInfo.IsDir() {
			files := []string{filepath.Join(templateDir, "layouts", "dashboard_layout.html"), page}
			files = append(files, commons...)
			renderer.AddFromFilesFuncs(filepath.Base(page), templateFuncMap, files...)
		}
	}

	return renderer
}

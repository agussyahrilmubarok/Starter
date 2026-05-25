package controller

import (
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func LoadTemplate(templateDir string) multitemplate.Renderer {
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

func render(c *gin.Context, code int, tmpl string, data gin.H) {
	c.HTML(code, tmpl, data)
}

package app

import (
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplate(templateDir string) multitemplate.Renderer {
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

	return renderer
}

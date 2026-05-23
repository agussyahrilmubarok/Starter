package app

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (app *App) setGinRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(app.requestIDMiddleware())
	router.Use(sessions.Sessions(app.cfg.App.Name, cookie.NewStore([]byte(app.cfg.App.Name))))

	router.HTMLRender = loadTemplate("./public/templates")
	router.Static("/static", "./public/static")
	router.StaticFile("/favicon.ico", "./public/static/favicon.ico")

	homeController := controller.NewHomeController()
	authController := controller.NewAuthController()

	router.GET("/", homeController.Index)

	router.GET("/sign-up", authController.SignUpPage)
	router.POST("/sign-up", authController.SignUp)

	return router
}

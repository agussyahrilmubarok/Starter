package application

import (
	"net/http"

	"agussyahrilmubarok.github.io/web/internal/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (app *App) newGinRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(app.requestIDMiddleware())

	store := cookie.NewStore([]byte(app.cfg.Session.SecretKey))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   app.cfg.Session.MaxAge,
		HttpOnly: true,
		// Secure: true, // (HTTPS)
	})
	router.Use(sessions.Sessions(app.cfg.App.Name, store))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "not_found.html", nil)
	})

	router.HTMLRender = controller.LoadTemplate("./public/templates")
	router.Static("/static", "./public/static")
	router.StaticFile("/favicon.ico", "./public/static/favicon.ico")

	homeController := controller.NewHomeController()
	authController := controller.NewAuthController(app.authService)
	dashboardController := controller.NewDashboardController()
	profileController := controller.NewProfileController()
	userController := controller.NewUserController(app.userService)

	public := router.Group("/")
	{
		public.GET("/", homeController.Index)
	}

	guest := router.Group("/")
	guest.Use(app.guestMiddleware())
	{
		guest.GET("/sign-up", authController.SignUp)
		guest.POST("/sign-up", authController.SignUp)
		guest.GET("/sign-in", authController.SignIn)
		guest.POST("/sign-in", authController.SignIn)
	}

	private := router.Group("/")
	private.Use(app.authMiddleware())
	{
		private.GET("/dashboard", dashboardController.Dashboard)
		private.POST("/sign-out", authController.SignOut)

		profile := private.Group("/dashboard/profile")
		{
			profile.GET("", profileController.Index)
		}

		users := private.Group("/dashboard/users")
		{
			users.GET("", userController.Index)
			users.GET("/create", userController.Create)
			users.POST("/create", userController.Store)
			users.GET("/:id/edit", userController.Edit)
			users.POST("/:id/edit", userController.Update)
			users.POST("/:id/delete", userController.Delete)
		}
	}

	return router
}

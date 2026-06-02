package web

import (
	"agussyahrilmubarok.github.io/web/internal/delivery/web/controller"
	"agussyahrilmubarok.github.io/web/internal/delivery/web/middleware"
	"github.com/gin-gonic/gin"
)

func Register(
	r *gin.Engine,
	appController *controller.AppController,
) {
	r.Use(middleware.RequestIDMiddleware())

	// r.NoRoute(func(c *gin.Context) {
	// 	c.HTML(http.StatusNotFound, "not_found.html", gin.H{
	// 		"Title": "404",
	// 	})
	// })

	r.HTMLRender = appController.LoadTemplate("./public/templates")
	r.Static("/static", "./public/static")
	r.StaticFile("/favicon.ico", "./public/static/favicon.ico")

	public := r.Group("/")
	{
		public.GET("/", appController.HomePage)
	}

	guest := r.Group("/")
	{
		guest.GET("/sign-up", appController.SignUp)
		guest.POST("/sign-up", appController.SignUp)
		guest.GET("/sign-in", appController.SignIn)
		guest.POST("/sign-in", appController.SignIn)
	}

	private := r.Group("/")
	private.Use(middleware.AuthMiddleware())
	{
		private.GET("/dashboard", appController.DashboardPage)
		private.POST("/sign-out", appController.SignOut)

		profile := private.Group("/dashboard/profile")
		{
			profile.GET("", appController.ProfilePage)
		}

		users := private.Group("/dashboard/users")
		{
			users.GET("", appController.UserListPage)
			users.GET("/add", appController.UserAddPage)
			users.POST("/add", appController.UserAdd)
			users.GET("/:id/edit", appController.UserEditPage)
			users.POST("/:id/edit", appController.UserEdit)
			users.POST("/:id/delete", appController.UserDelete)
		}
	}
}

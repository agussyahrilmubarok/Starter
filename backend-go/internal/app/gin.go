package app

import (
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "agussyahrilmubarok.github.io/backend/api/gin/docs"
)

// @title Backend API
// @version 1.0
// @description Backend API
// @BasePath /api
func (app *App) setGinRouter() http.Handler {
	router := gin.Default()

	authHandler := handler.NewAuthHandler(app.authService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sign-up", authHandler.SignUp)
			auth.POST("/sign-in", authHandler.SignIn)
		}
	}

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName(app.config.App.Name),
		ginSwagger.URL("/api/swagger/doc.json"),
	))

	return router
}

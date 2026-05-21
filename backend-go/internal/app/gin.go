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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Input your Bearer token in this format: Bearer <token>
func (app *App) setGinRouter() http.Handler {
	router := gin.Default()

	authHandler := handler.NewAuthHandler(app.authService)
	userHandler := handler.NewUserHandler(app.userService)

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

		users := v1.Group("/users")
		users.Use(app.requireAuthMiddleware())
		{
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.POST("", userHandler.Create)
			users.PUT("/:id", userHandler.UpdateByID)
			users.DELETE("/:id", userHandler.DeleteByID)
		}
	}

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName(app.config.App.Name),
		ginSwagger.URL("/api/swagger/doc.json"),
	))

	return router
}

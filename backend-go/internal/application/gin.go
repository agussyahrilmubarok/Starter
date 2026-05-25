package application

import (
	"net/http"

	"agussyahrilmubarok.github.io/backend/internal/handler"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "agussyahrilmubarok.github.io/backend/api/docs"
)

// @title           Backend REST API
// @version         1.0.0
// @description     RESTful API documentation for backend services.
// @BasePath        /api

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Enter JWT Bearer token in the format: Bearer <token>
func (app *App) newGinRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	authHandler := handler.NewAuthHandler(app.authService)
	userHandler := handler.NewUserHandler(app.userService)

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

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName(app.cfg.App.Name),
		ginSwagger.URL("/swagger-ui/doc.json"),
	))

	return router
}

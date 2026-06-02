package http

import (
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/handler"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/middleware"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"github.com/gin-gonic/gin"

	_ "agussyahrilmubarok.github.io/backend/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Register(
	r *gin.Engine,
	jwtManager security.JWTManager,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.RequestIDMiddleware())

	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sign-up", authHandler.SignUp)
			auth.POST("/sign-in", authHandler.SignIn)
		}

		users := v1.Group("/users", middleware.Auth(jwtManager))
		{
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.POST("", userHandler.Create)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}
}

func RegisterSwaggerUI(r *gin.Engine) {
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName("swagger"),
		ginSwagger.URL("/swagger-ui/doc.json"),
	))
}

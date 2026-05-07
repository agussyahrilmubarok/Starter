package main

import (
	"backend/internal/common"
	"backend/internal/common/middlewares"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/config"
	"backend/pkg/database"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	gin.SetMode(gin.ReleaseMode)

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("[ERROR] failed to connect database err:%v\n", err)
	}

	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.GET("/health", common.HealthHandler)

	apiV2 := router.Group("/api/v2")
	{
		apiV2.POST("/register", authHandler.Register)
		apiV2.POST("/login", authHandler.Login)

		apiV2.GET("/users", middlewares.AuthMiddleware(), userHandler.FindUsers)
		apiV2.GET("/users/:id", middlewares.AuthMiddleware(), userHandler.FindUserById)
		apiV2.POST("/users", middlewares.AuthMiddleware(), userHandler.CreateUser)
		apiV2.PUT("/users/:id", middlewares.AuthMiddleware(), userHandler.UpdateUser)
		apiV2.DELETE("/users/:id", middlewares.AuthMiddleware(), userHandler.DeleteUser)
	}

	router.Run(fmt.Sprintf(":%v", config.GetEnv("APP_PORT", "3000")))
}

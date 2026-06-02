package main

import (
	"log"

	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	deliveryhttp "agussyahrilmubarok.github.io/backend/internal/delivery/http"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/handler"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/repository/postgres"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/server"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"go.uber.org/zap"
)

// @title           Backend REST API
// @version         1.0.0
// @description     RESTful API documentation for backend services.
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Enter your JWT token in the format: Bearer {token}
func main() {
	cfg, err := config.Load("configs/config.yml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := logger.Init(cfg.App.Logger.FilePath, logger.ParseLevel(cfg.App.Logger.Level)); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	db, err := config.NewPostgres(cfg)
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}

	jwtManager := security.NewJwtManager(&cfg.App.JWT)

	userRepo := postgres.NewUserRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo, jwtManager)
	userUC := usecase.NewUserUseCase(userRepo)

	authHandler := handler.NewAuthHandler(authUC)
	userHandler := handler.NewUserHandler(userUC)

	srv := server.NewHTTPServer(cfg)
	deliveryhttp.Register(srv.Router(), jwtManager, authHandler, userHandler)
	deliveryhttp.RegisterSwaggerUI(srv.Router())

	if err := srv.Run(); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}

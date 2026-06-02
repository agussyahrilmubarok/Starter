package main

import (
	"log"

	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/delivery/http/handler"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/repository/postgres"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/server"
	"agussyahrilmubarok.github.io/backend/pkg/logger"

	deliveryhttp "agussyahrilmubarok.github.io/backend/internal/delivery/http"
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
		log.Fatalf("failed to connect database: %v", err)
	}

	jwtManager := security.NewJwtManager(&cfg.App.JWT)

	userRepo := postgres.NewUserRepository(db)

	authUC := usecase.NewAuthUseCase(userRepo, jwtManager)
	userUC := usecase.NewUserUseCase(userRepo)

	authHandler := handler.NewAuthHandler(authUC)
	userHandler := handler.NewUserHandler(userUC)

	srv := server.NewHTTPServer(cfg)
	deliveryhttp.Register(srv.Router(), cfg, jwtManager, authHandler, userHandler)
	deliveryhttp.RegisterSwaggerUI(srv.Router())

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

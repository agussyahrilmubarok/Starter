package main

import (
	"log"

	"agussyahrilmubarok.github.io/web/internal/delivery/web/controller"
	"agussyahrilmubarok.github.io/web/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/web/internal/infrastructure/repository/postgres"
	"agussyahrilmubarok.github.io/web/internal/infrastructure/server"
	"agussyahrilmubarok.github.io/web/pkg/logger"

	deliveryweb "agussyahrilmubarok.github.io/web/internal/delivery/web"
)

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

	userRepo := postgres.NewUserRepository(db)

	appController := controller.NewAppController(userRepo)

	srv := server.NewWEBServer(cfg)
	deliveryweb.Register(srv.Router(), appController)

	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

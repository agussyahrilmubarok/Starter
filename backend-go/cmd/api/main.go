package main

import (
	"log"

	"agussyahrilmubarok.github.io/backend/internal/app"
	"agussyahrilmubarok.github.io/backend/internal/config"
	"agussyahrilmubarok.github.io/backend/pkg/database"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	if err := logger.Init("logs/app.log"); err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Log.Sync()

	cfg, err := config.Load("configs/config.yml")
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	db, err := database.NewPostgres(&database.PostgresConfig{
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		Port:     cfg.Database.Port,
		SSLMode:  cfg.Database.SSLMode,
		TimeZone: cfg.Database.TimeZone,
	})
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	app := app.NewApp(cfg, db)
	if err := app.Run(); err != nil {
		logger.Fatal("Failed to run application", zap.Error(err))
	}
}

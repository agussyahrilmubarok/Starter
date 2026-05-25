package main

import (
	"log"

	"agussyahrilmubarok.github.io/backend/internal/application"
	"agussyahrilmubarok.github.io/backend/internal/config"
	"agussyahrilmubarok.github.io/backend/pkg/database"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
)

func main() {
	cfg, err := config.Load("configs/config.yml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = logger.Init(logger.DefaultOptions())
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	db, err := database.NewPostgres(database.PostgresConfig{
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		Port:     cfg.Database.Port,
		SSLMode:  cfg.Database.SSLMode,
		TimeZone: cfg.Database.TimeZone,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	app := application.New(cfg, db)
	if err := app.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}

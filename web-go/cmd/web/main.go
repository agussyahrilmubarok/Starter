package main

import (
	"log"

	"agussyahrilmubarok.github.io/web/internal/app"
	"agussyahrilmubarok.github.io/web/internal/config"
	"agussyahrilmubarok.github.io/web/pkg/database"
	"agussyahrilmubarok.github.io/web/pkg/logger"
)

func main() {
	cfg, err := config.Load("configs/config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	err = logger.Init(logger.DefaultOptions())
	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}
	defer logger.Sync()

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
		log.Fatalf("Failed to connect to database: %v", err)
	}

	application := app.NewApp(cfg, db)
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}

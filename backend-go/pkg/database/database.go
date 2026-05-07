package database

import (
	"backend/internal/entity"
	"backend/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] failed to connect to database: %v\n", err)
		return nil, err
	}
	log.Printf("[INFO] database connected successfully")

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("[ERROR] failed to migrated database: %v\n", err)
		return nil, err
	}
	log.Printf("[INFO] database migrated successfully\n")

	return db, nil
}

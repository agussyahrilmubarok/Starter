package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresConfig stores database credentials and settings.
type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

// NewPostgres creates a new PostgreSQL connection using GORM.
func NewPostgres(cfg PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Set logger to Info for development, change to Error/Silent for production
		Logger: logger.Default.LogMode(logger.Silent),
		// Optional: Disable auto-ping during initialization for a slightly faster startup,
		DisableAutomaticPing: false,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns: The maximum number of connections in the idle connection pool.
	// Prevents the app from having to establish a new connection from scratch for every request.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns: The maximum number of open connections to the database (idle + in use).
	// Adjust this based on your database server's RAM and connection limits.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime: The maximum amount of time a connection may be reused.
	// Crucial to prevent connections from going stale and being abruptly dropped by firewalls/networks.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Name            string `mapstructure:"name"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	SSLMode         string `mapstructure:"ssl_mode"`
	Timezone        string `mapstructure:"timezone"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

func NewPostgres(cfg *Config) (*gorm.DB, error) {
	dbCfg := cfg.Postgres

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.SSLMode,
		dbCfg.Timezone,
	)

	gormCfg := &gorm.Config{}

	if cfg.App.Env != "production" {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnMaxLifetime) * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return db, nil
}

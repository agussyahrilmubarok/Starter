package config

import (
	"strings"

	"agussyahrilmubarok.github.io/backend/pkg/database"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database database.PostgresConfig
	JWT      JWTConfig
}

// Load config file, filepath = "configs/config.yml"
func Load(filepath string) (*Config, error) {
	lastSlash := strings.LastIndex(filepath, "/")
	lastDot := strings.LastIndex(filepath, ".")

	dir := filepath[:lastSlash+1]
	filename := filepath[lastSlash+1 : lastDot]
	ext := filepath[lastDot+1:]

	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Name:     viper.GetString("app.name"),
			Port:     viper.GetString("app.port"),
			TimeZone: viper.GetString("app.timezone"),
		},
		Database: database.PostgresConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
		JWT: JWTConfig{
			SecretKey: viper.GetString("jwt.secret_key"),
			ExpTime:   viper.GetString("jwt.exp_time"),
		},
	}, nil
}

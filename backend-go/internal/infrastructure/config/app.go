package config

type App struct {
	Name   string `mapstructure:"name"`
	Port   int    `mapstructure:"port"`
	Env    string `mapstructure:"env"`
	JWT    JWT    `mapstructure:"jwt"`
	Logger Logger `mapstructure:"logger"`
	CORS   CORS   `mapstructure:"cors"`
}

type JWT struct {
	Secret     string `mapstructure:"secret"`
	ExpiryHour int    `mapstructure:"expiry_hour"`
}

type Logger struct {
	FilePath string `mapstructure:"file_path"`
	Level    string `mapstructure:"level"`
}

type CORS struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

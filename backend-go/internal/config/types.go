package config

type AppConfig struct {
	Name     string
	Port     string
	TimeZone string
}

type JWTConfig struct {
	SecretKey string
	ExpTime   string
}

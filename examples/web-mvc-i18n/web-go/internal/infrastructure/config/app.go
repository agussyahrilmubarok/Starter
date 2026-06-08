package config

type App struct {
	Name    string  `mapstructure:"name"`
	Port    int     `mapstructure:"port"`
	Env     string  `mapstructure:"env"`
	Session Session `mapstructure:"session"`
	Logger  Logger  `mapstructure:"logger"`
}

type Session struct {
	Secret string `mapstructure:"secret"`
	MaxAge int    `mapstructure:"max_age"`
}

type Logger struct {
	FilePath string `mapstructure:"file_path"`
	Level    string `mapstructure:"level"`
}

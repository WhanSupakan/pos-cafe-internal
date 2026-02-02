package config

// Application configuration structs
// This file will contain all configuration structs
type Config struct {
	Port        string `mapstructure:"port"`
	Env         string `mapstructure:"env"`
	DatabaseDSN string `mapstructure:"db_dsn"`
}

func NewConfig() *Config {
	return &Config{
		Port:        "8080",
		Env:         "development",
		DatabaseDSN: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	}
}

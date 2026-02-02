package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Load configuration from environment variables
// This file will contain the configuration loading logic
func LoadConfig() *Config {
	cfg := NewConfig()

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read .env file: %v. Using environment variables or defaults.", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return cfg
}

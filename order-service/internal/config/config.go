package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUrl string
	HttpPort    string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // look for config in the working directory
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)

	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	cfg.DatabaseUrl = os.Getenv("DATABASE_URL")

	return &cfg, nil
}

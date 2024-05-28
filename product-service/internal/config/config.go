package config

import (
	"fmt"
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

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)

	}
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	return &cfg, nil
}

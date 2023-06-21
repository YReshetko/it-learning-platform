package config

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/config"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	DB         config.Database
	GRPCServer config.GRPCServer
}

func LoadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to load app initial config: %w", err)
	}
	return cfg, nil
}

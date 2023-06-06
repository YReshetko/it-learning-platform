package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	HTTP       HTTP
	AuthClient AuthClient
}

type HTTP struct {
	Host string `env:"HTTP_HOST"`
	Port int    `env:"HTTP_PORT"`
}

type AuthClient struct {
	Host string `env:"AUTH_HOST"`
	Port int    `env:"AUTH_PORT"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to load app initial config: %w", err)
	}
	return cfg, nil
}

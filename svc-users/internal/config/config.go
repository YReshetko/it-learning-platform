package config

import (
	"encoding/json"
	"fmt"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/config"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	DB         Database
	GRPCServer config.GRPCServer
}

type Database struct {
	Host         string `env:"DB_HOST"`
	Port         int    `env:"DB_PORT"`
	UserName     string `env:"DB_USER_NAME"`
	UserPassword string `env:"DB_USER_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
}

type GRPC struct {
	Host    string `env:"GRPC_HOST"`
	Port    int    `env:"GRPC_PORT"`
	Network string `env:"GRPC_NETWORK"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to load app initial config: %w", err)
	}
	return cfg, nil
}

func (c Config) ToJSON() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("unable to marshal config to JSON: %w", err)
	}
	return string(b), nil
}

func (db Database) CreatePostgresDSN() string {
	template := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	return fmt.Sprintf(template, db.Host, db.UserName, db.UserPassword, db.DatabaseName, db.Port)
}

package config

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/config"
	"github.com/caarlos0/env/v8"
	_ "github.com/caarlos0/env/v8"
)

type Config struct {
	UsersClient    UsersClient
	KeycloakClient KeycloakClient
	GRPCServer     config.GRPCServer
}

type UsersClient struct {
	Host string `env:"USERS_HOST"`
	Port int    `env:"USERS_PORT"`
}

type KeycloakClient struct {
	Host         string `env:"KEYCLOAK_HOST"`
	Port         int    `env:"KEYCLOAK_PORT"`
	Realm        string `env:"KEYCLOAK_REALM"`
	ClientID     string `env:"KEYCLOAK_CLIENT_ID"`
	ClientSecret string `env:"KEYCLOAK_CLIENT_SECRET"`
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

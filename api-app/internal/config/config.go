package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	HTTP          HTTP
	AuthClient    AuthClient
	AuthRedirect  AuthRedirect
	CoursesClient CoursesClient
}

type HTTP struct {
	Host string `env:"HTTP_HOST"`
	Port int    `env:"HTTP_PORT"`
}

type AuthClient struct {
	Host string `env:"AUTH_HOST"`
	Port int    `env:"AUTH_PORT"`
}

type CoursesClient struct {
	Host string `env:"COURSES_HOST"`
	Port int    `env:"COURSES_PORT"`
}

// "http://localhost:8081/realms/it-academy/protocol/openid-connect/auth?response_type=token&scope=openid&client_id=academy&redirect_uri=http://localhost:8080"
const redirectTemplate = "%s://%s:%d/realms/%s/protocol/openid-connect/auth?response_type=token&scope=openid&client_id=academy&redirect_uri=%s"

type AuthRedirect struct {
	Schema      string `env:"AUTH_REDIRECT_SCHEMA"`
	Host        string `env:"AUTH_REDIRECT_HOST"`
	Port        int    `env:"AUTH_REDIRECT_PORT"`
	Realm       string `env:"AUTH_REDIRECT_REALM"`
	RedirectURI string `env:"AUTH_REDIRECT_URI"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to load app initial config: %w", err)
	}
	return cfg, nil
}

func (ar AuthRedirect) URL() string {
	return fmt.Sprintf(redirectTemplate, ar.Schema, ar.Host, ar.Port, ar.Realm, ar.RedirectURI)
}

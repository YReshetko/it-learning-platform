package config

import "fmt"

type GRPCServer struct {
	Host    string `env:"GRPC_HOST"`
	Port    int    `env:"GRPC_PORT"`
	Network string `env:"GRPC_NETWORK"`
}

type Database struct {
	Host         string `env:"DB_HOST"`
	Port         int    `env:"DB_PORT"`
	UserName     string `env:"DB_USER_NAME"`
	UserPassword string `env:"DB_USER_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
}

func (db Database) CreatePostgresDSN() string {
	template := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	return fmt.Sprintf(template, db.Host, db.UserName, db.UserPassword, db.DatabaseName, db.Port)
}

package main

import (
	"github.com/YReshetko/it-academy-cources/svc-users/internal/config"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/grpc"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/storage"
	"log"
)

func main() {
	cfg := handleError(config.LoadConfig())
	db := handleError(storage.DatabaseConnection(cfg.DB))
	s := storage.NewUserStorage(db)

	server := grpc.NewServer(cfg.GRPC, grpc.NewHandler(s))
	server.Start()
}

func handleError[T any](val T, err error) T {
	if err != nil {
		log.Fatalf("crashing app due to: %s", err)
	}
	return val
}

package main

import (
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/clients"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/grpc"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	logger := logrus.New().WithField("application", "svc-users")

	cfg := handleError(config.LoadConfig())
	uc := clients.NewUsersClient(cfg.UsersClient)
	kc := clients.NewKeycloakClient(cfg.KeycloakClient)
	h := grpc.NewHandler(&kc, &uc)
	server := grpc.NewServer(cfg.GRPC, &h)
	server.Start()
}

func handleError[T any](val T, err error) T {
	if err != nil {
		log.Fatalf("crashing app due to: %s", err)
	}
	return val
}

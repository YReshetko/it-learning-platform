package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	commonGrpc "github.com/YReshetko/it-learning-platform/lib-app/pkg/grpc"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/clients"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/grpc"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.StandardLogger().WithField("application", "svc-auth")

	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	uc := clients.NewUsersClient(cfg.UsersClient, logger.WithField("client", "UsersClient"))
	kc := clients.NewKeycloakClient(cfg.KeycloakClient)
	h := grpc.NewHandler(&kc, logger.WithField("handler", "grpc"), &uc)
	server := commonGrpc.NewServer[auth.AuthServiceServer](
		commonGrpc.WithCfg[auth.AuthServiceServer](cfg.GRPCServer),
		commonGrpc.WithHandler[auth.AuthServiceServer](&h),
		commonGrpc.WithRegistrarFn[auth.AuthServiceServer](auth.RegisterAuthServiceServer),
		commonGrpc.WithLogger[auth.AuthServiceServer](logger.WithField("server", "grpc")),
	)
	server.Start()
}

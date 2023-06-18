package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	commonGrpc "github.com/YReshetko/it-learning-platform/lib-app/pkg/grpc"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/grpc"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.StandardLogger().WithField("application", "svc-users")

	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	db := errors.MustExitAppErrorHandler[*gorm.DB](logger.WithField("sub_system", "database"))(storage.DatabaseConnection(cfg.DB))
	s := storage.NewUserStorage(db)

	handler := grpc.NewHandler(logger.WithField("handler", "grpc"), s)

	server := commonGrpc.NewServer[users.UserServiceServer](
		commonGrpc.WithCfg[users.UserServiceServer](cfg.GRPCServer),
		commonGrpc.WithHandler[users.UserServiceServer](&handler),
		commonGrpc.WithRegistrarFn[users.UserServiceServer](users.RegisterUserServiceServer),
		commonGrpc.WithLogger[users.UserServiceServer](logger.WithField("server", "grpc")),
	)

	server.Start()
}

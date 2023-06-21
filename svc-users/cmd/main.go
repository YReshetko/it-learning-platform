package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/db"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	libGrpc "github.com/YReshetko/it-learning-platform/lib-app/pkg/grpc"
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
	dbConnection := errors.MustExitAppErrorHandler[*gorm.DB](logger.WithField("sub_system", "database"))(db.DatabaseConnection(cfg.DB))
	s := storage.NewUserStorage(dbConnection)

	handler := grpc.NewHandler(logger.WithField("handler", "grpc"), s)

	server := libGrpc.NewServer[users.UserServiceServer](
		libGrpc.WithCfg[users.UserServiceServer](cfg.GRPCServer),
		libGrpc.WithHandler[users.UserServiceServer](&handler),
		libGrpc.WithRegistrarFn[users.UserServiceServer](users.RegisterUserServiceServer),
		libGrpc.WithLogger[users.UserServiceServer](logger.WithField("server", "grpc")),
	)

	server.Start()
}

package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/grpc"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New().WithField("application", "svc-users")

	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	db := errors.MustExitAppErrorHandler[*gorm.DB](logger.WithField("sub_system", "database"))(storage.DatabaseConnection(cfg.DB))
	s := storage.NewUserStorage(db)

	handler := grpc.NewHandler(logger.WithField("handler", "grpc"), s)
	server := grpc.NewServer(cfg.GRPC, &handler, logger.WithField("server", "grpc"))
	server.Start()
}

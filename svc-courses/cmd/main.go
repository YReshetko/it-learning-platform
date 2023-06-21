package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/db"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.StandardLogger().WithField("application", "svc-courses")
	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	dbConnection := errors.MustExitAppErrorHandler[*gorm.DB](logger.WithField("sub_system", "database"))(db.DatabaseConnection(cfg.DB))

	_ = storage.NewCourseStorage(dbConnection)
}

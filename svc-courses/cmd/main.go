package main

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/db"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	libGrpc "github.com/YReshetko/it-learning-platform/lib-app/pkg/grpc"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/grpc"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.StandardLogger().WithField("application", "svc-courses")
	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	dbConnection := errors.MustExitAppErrorHandler[*gorm.DB](logger.WithField("sub_system", "database"))(db.DatabaseConnection(cfg.DB))

	s := storage.NewCourseStorage(dbConnection)
	h := grpc.NewHandler(logger.WithField("handler", "grpc"), s, grpc.TechnologyMapperImpl{})

	server := libGrpc.NewServer[courses.CoursesServiceServer](
		libGrpc.WithCfg[courses.CoursesServiceServer](cfg.GRPCServer),
		libGrpc.WithHandler[courses.CoursesServiceServer](&h),
		libGrpc.WithRegistrarFn[courses.CoursesServiceServer](courses.RegisterCoursesServiceServer),
		libGrpc.WithLogger[courses.CoursesServiceServer](logger.WithField("server", "grpc")),
	)

	server.Start()
}

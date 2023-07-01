package main

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/clients"
	"github.com/YReshetko/it-learning-platform/api-app/internal/config"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/handlers"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/mappers"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/authorization"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/routes"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.StandardLogger().WithField("application", "api-app")

	cfg := errors.MustExitAppErrorHandler[config.Config](logger.WithField("sub_system", "config"))(config.LoadConfig())
	server := http.NewServer(cfg.HTTP, logger.WithField("server", "http"))

	authClient := clients.NewAuthClient(cfg.AuthClient, logger.WithField("client", "AuthClient"))
	coursesClient := clients.NewCoursesClient(cfg.CoursesClient, logger.WithField("client", "CoursesClient"))

	registration := handlers.NewRegistration(authClient, logger.WithField("handler", "Registration"))
	self := handlers.NewSelf(authClient, logger.WithField("handler", "Self"))
	courses := handlers.NewCourses(
		handlers.WithClient(coursesClient),
		handlers.WithLogger(logger.WithField("handler", "Courses")),
		handlers.WithTechnologyMapper(mappers.TechnologyMapperImpl{}),
		handlers.WithCategoryMapper(mappers.CategoryMapperImpl{}),
		handlers.WithTopicMapper(mappers.TopicMapperImpl{}),
	)

	authorizationService := authorization.NewService(authClient)

	routeServices := routes.NewRouterServices(
		routes.WithLogger(logger.WithField("sub_system", "RouterServices")),
		routes.WithAuthService(&authorizationService),
		routes.WithRedirectURL(cfg.AuthRedirect.URL()),
	)
	r := routes.NewRouter(
		routes.WithServices(&routeServices),
		routes.WithRegistration(registration),
		routes.WithSelf(self),
		routes.WithCourses(courses),
	)
	r.Init(server.Engine)
	server.Start()

}

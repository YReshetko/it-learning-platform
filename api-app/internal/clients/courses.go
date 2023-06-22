package clients

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
CoursesClient the client to RGPC svc-courses API
@Constructor
*/
type CoursesClient struct {
	cfg                          config.CoursesClient
	logger                       *logrus.Entry
	courses.CoursesServiceClient // @Exclude
}

// @PostConstruct
func (uc *CoursesClient) postConstruct() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	host := fmt.Sprintf("%s:%d", uc.cfg.Host, uc.cfg.Port)
	conn, err := grpc.Dial(host, opt...)
	if err != nil {
		uc.logger.WithField("host", host).WithError(err).Error("Unable to establish deal connection")
		return
	}
	uc.CoursesServiceClient = courses.NewCoursesServiceClient(conn)
}

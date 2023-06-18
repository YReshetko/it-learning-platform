package clients

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

/*
UsersClient the client to RGPC svc-users API
@Constructor
*/
type UsersClient struct {
	cfg    config.UsersClient
	logger *logrus.Entry

	users.UserServiceClient // @Exclude
}

// @PostConstruct
func (uc *UsersClient) postConstruct() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", uc.cfg.Host, uc.cfg.Port), opt...)
	if err != nil {
		uc.logger.WithError(err).Error("Unable to set up Dial connection")
		os.Exit(2)
	}
	uc.UserServiceClient = users.NewUserServiceClient(conn)
}

package clients

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
AuthClient the client to RGPC svc-users API
@Constructor
*/
type AuthClient struct {
	cfg                    config.AuthClient
	logger                 *logrus.Entry
	auth.AuthServiceClient // @Exclude
}

// @PostConstruct
func (uc *AuthClient) postConstruct() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	host := fmt.Sprintf("%s:%d", uc.cfg.Host, uc.cfg.Port)
	conn, err := grpc.Dial(host, opt...)
	if err != nil {
		uc.logger.WithField("host", host).WithError(err).Error("Unable to establish deal connection")
		return
	}
	uc.AuthServiceClient = auth.NewAuthServiceClient(conn)
}

package clients

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/*
AuthClient the client to RGPC svc-users API
@Constructor
*/
type AuthClient struct {
	cfg                    config.AuthClient
	auth.AuthServiceClient // @Exclude
}

// @PostConstruct
func (uc *AuthClient) postConstruct() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", uc.cfg.Host, uc.cfg.Port), opt...)
	if err != nil {
		log.Fatalln(err)
	}
	uc.AuthServiceClient = auth.NewAuthServiceClient(conn)
}

package clients

import (
	"fmt"
	"github.com/YReshetko/it-academy-cources/svc-auth/internal/config"
	"github.com/YReshetko/it-academy-cources/svc-users/pb/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/*
UsersClient the client to RGPC svc-users API
@Constructor
*/
type UsersClient struct {
	cfg                     config.UsersClient
	users.UserServiceClient // @Exclude
}

// @PostConstruct
func (uc *UsersClient) postConstruct() {
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", uc.cfg.Host, uc.cfg.Port), opt...)
	if err != nil {
		log.Fatalln(err)
	}
	uc.UserServiceClient = users.NewUserServiceClient(conn)
}

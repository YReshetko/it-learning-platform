package grpc

import (
	"fmt"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/config"
	"github.com/YReshetko/it-academy-cources/svc-users/pb/users"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
Server the GRPC server handler
@Constructor
*/
type Server struct {
	cfg     config.GRPC
	handler Handler

	server *grpc.Server // @Exclude
}

func (s Server) Start() {
	server := grpc.NewServer()
	users.RegisterUserServiceServer(server, s.handler)
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	fmt.Println("Listening:", addr)
	listener, err := net.Listen(s.cfg.Network, addr)
	if err != nil {
		log.Fatalln(err)
	}
	s.server = server
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s Server) Stop() {
	s.server.GracefulStop()
}

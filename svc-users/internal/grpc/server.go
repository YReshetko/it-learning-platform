package grpc

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/config"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

/*
Server the GRPC server handler
@Constructor
*/
type Server struct {
	cfg     config.GRPC
	handler *Handler
	logger  *logrus.Entry

	server *grpc.Server // @Exclude
}

func (s Server) Start() {
	server := grpc.NewServer()
	users.RegisterUserServiceServer(server, s.handler)
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	logger := s.logger.WithField("address", addr)
	listener, err := net.Listen(s.cfg.Network, addr)
	if err != nil {
		logger.WithError(err).Error("Unable to init listener")
		return
	}
	s.server = server
	logger.Info("Start listening")
	err = server.Serve(listener)
	if err != nil {
		logger.WithError(err).Error("Closed unexpectedly")
		return
	}
	logger.Info("Stopped gracefully")
}

func (s Server) Stop() {
	s.server.GracefulStop()
}

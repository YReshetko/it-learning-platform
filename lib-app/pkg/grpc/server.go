package grpc

import (
	"context"
	"fmt"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

/*
Server the generic GRPC server handler
T the server handler type
@Optional
*/
type Server[T any] struct {
	cfg         config.GRPCServer
	registrarFn func(grpc.ServiceRegistrar, T)
	handler     T
	logger      *logrus.Entry

	server   *grpc.Server // @Exclude
	listener net.Listener // @Exclude
}

func (s *Server[T]) Start() {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(contextPropagationUnaryServerInterceptor()),
	}

	server := grpc.NewServer(opts...)
	s.registrarFn(server, s.handler)

	addr := fmt.Sprintf(":%d", s.cfg.Port)
	s.logger = s.logger.WithField("address", addr)
	listener, err := net.Listen(s.cfg.Network, addr)

	if err != nil {
		s.logger.WithError(err).Error("Unable to init listener")
		return
	}
	s.server = server
	s.listener = listener

	s.logger.Info("Start listening")
	err = s.server.Serve(s.listener)
	if err != nil {
		s.logger.WithError(err).Error("Closed unexpectedly")
		return
	}
	s.logger.Info("Stopped gracefully")
}

func (s *Server[T]) Stop() {
	s.server.GracefulStop()
	s.logger.Info("Server stopped")
}

func contextPropagationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return handler(ctx, req)
	}
}

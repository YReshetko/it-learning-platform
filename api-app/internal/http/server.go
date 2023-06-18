package http

import (
	"context"
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

/*
Server the HTTP server for public API
@Constructor
*/
type Server struct {
	cfg    config.HTTP
	logger *logrus.Entry

	server *http.Server // @Exclude
	Engine *gin.Engine  // @Exclude
}

// @PostConstruct
func (s *Server) postConstruct() {
	s.Engine = gin.Default()
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	s.logger = s.logger.WithField("address", addr)
	s.server = &http.Server{
		Addr:                         addr,
		Handler:                      s.Engine,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
	}
}

func (s *Server) Start() {
	s.logger.Info("Starting server")
	err := s.server.ListenAndServe()
	if err != nil {
		s.logger.WithError(err).Error("Server closed")
		return
	}
	s.logger.Info("Server stopped gracefully")
}

func (s *Server) Stop() {
	s.logger.Info("Stopping server")
	err := s.server.Shutdown(context.Background())
	if err != nil {
		s.logger.WithError(err).Error("Server closed")
		return
	}
	s.logger.Info("Server stopped gracefully")
}

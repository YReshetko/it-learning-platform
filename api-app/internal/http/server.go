package http

import (
	"context"
	"fmt"
	"github.com/YReshetko/it-academy-cources/api-app/internal/config"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/*
Server the HTTP server for public API
@Constructor
*/
type Server struct {
	cfg config.HTTP

	server *http.Server // @Exclude
	Engine *gin.Engine  // @Exclude
}

// @PostConstruct
func (s *Server) postConstruct() {
	s.Engine = gin.Default()

	s.server = &http.Server{
		Addr:                         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:                      s.Engine,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
	}
}

func (s *Server) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) Stop() {
	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}

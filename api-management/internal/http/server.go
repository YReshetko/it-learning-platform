package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	registerRoutes(s.router)

	return s
}

func (s *Server) Start() {
	http.ListenAndServe(":8000", s.router)
}

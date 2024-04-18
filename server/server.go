package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type server struct {
	Router chi.Router
	Port   string
}

func New() *server {
	return &server{
		Router: chi.NewRouter(),
		Port:   "8000",
	}
}

func (s *server) WithPort(port string) *server {
	s.Port = port
	return &server{}
}

func (s *server) Run() error {
	s.mapRoute()

	log.Printf("Listening on port :%s\n", s.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.Router)
}

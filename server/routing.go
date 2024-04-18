package server

import (
	"blogx/controller"
	"net/http"
)

func (s *server) mapRoute() {
	// STATIC
	s.Router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s.Router.Get("/dashboard/login", controller.LoginView())
	s.Router.Post("/dashboard/login", controller.Login())
	s.Router.Get("/dashboard", controller.DashboardView())
}

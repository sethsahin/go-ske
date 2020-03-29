package controllers

import (
	"skeleton/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}

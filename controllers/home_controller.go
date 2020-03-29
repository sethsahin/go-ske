package controllers

import (
	"net/http"
	"skeleton/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hello World")
}

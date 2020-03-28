package controllers

import (
	"net/http"
	"skeleton/db"
	"skeleton/utils"
)

func (server *Server) Home(w http.ResponseWriter, r http.Request) {
	responses.JSON(w, http.StatusOK, "Hello World")
}

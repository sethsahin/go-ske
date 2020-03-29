package controllers

import (
	"net/http"
	"skeleton/utils"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Hello World")
}

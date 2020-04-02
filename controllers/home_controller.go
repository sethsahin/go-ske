package controllers

import (
	"net/http"

	"skeleton/repositories"
	"skeleton/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	repo, err := repositories.GetAllPosts(server.DB)

	if err != nil {
		return
	}

	responses.JSON(w, http.StatusOK, repo)
}

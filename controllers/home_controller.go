package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"skeleton/models"
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

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postCreated, err := repositories.CreatePost(server.DB, &post)
	if err != nil {
		return
	}

	responses.JSON(w, http.StatusCreated, postCreated)
}

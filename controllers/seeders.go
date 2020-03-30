package controllers

import (
	"github.com/Pallinder/go-randomdata"
	"log"
	"math/rand"
	"skeleton/database/seeds"
	"skeleton/models"
)

func (server *Server) DatabaseSeeds() {
	err := server.DB.Debug().DropTableIfExists(&models.Post{}).Error
	if err != nil {
		log.Fatalln("Cannot drop tables, %v", err)
	}

	counter := 0

	for counter < 10 {
		seeds.CreatePost(server.DB, randomdata.FirstName(randomdata.Male), randomdata.FirstName(randomdata.Male), rand.Intn(50))
		counter += 1
	}

}

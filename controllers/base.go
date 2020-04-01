package controllers

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Mysql Database Driver
	"log"
	"math/rand"
	"net/http"
	"skeleton/database/seeds"

	"skeleton/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Connect(host, user, password, dbName, driver, port string) {

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)

	fmt.Println(DBURL)

	var err error

	server.DB, err = gorm.Open(driver, DBURL)

	if err != nil {
		log.Fatalf("%v, Connect error", err)
	} else {
		fmt.Println("%v, Connected", driver)
	}

	/**
	 *	  Migrations
	 */
	server.DB.Debug().AutoMigrate(models.Post{})

	/**
	 *    Seeds
	 */
	server.DatabaseSeeds()

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening port on 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

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
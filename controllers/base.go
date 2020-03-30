package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Mysql Database Driver
	"log"
	"net/http"

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

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening port on 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

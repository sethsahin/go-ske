package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Connect(host, user, password, dbName, driver, port string) {

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)

	if driver == "mysql" {
		db, err := gorm.Open(driver, DBURL)
		if err != nil {
			log.Fatalf("Mysql connect error, %v", err)
		}
		defer db.Close()
	} else if driver == "postgres" {
		db, err := gorm.Open(driver, DBURL)
		if err != nil {
			log.Fatalf("PostgreSQL connect error, %v", err)
		}
		defer db.Close()
	}
}

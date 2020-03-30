package seeds

import (
	"github.com/jinzhu/gorm"
	"skeleton/models"

	"log"
)

func CreatePost(db *gorm.DB, title, description string, order int) error {
	err := db.DropTableIfExists(&models.Post{}).Error
	if err != nil {
		log.Fatalln("%v", err)
	}

	err = db.Debug().AutoMigrate(&models.Post{}).Error
	if err != nil {
		log.Fatalln("%v", err)
	}

	return db.Create(&models.Post{Title: title, Description: description, Order: order}).Error
}

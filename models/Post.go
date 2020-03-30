package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)"`
	Description string
	Order       int `gorm:AUTO_INCREMENT`
}

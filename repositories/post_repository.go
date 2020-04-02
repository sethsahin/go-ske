package repositories

import (
	"github.com/jinzhu/gorm"
	"skeleton/models"
)

type PostRepository interface {
	GetAllPosts(db *gorm.DB) (*[]models.Post, error)
}

func GetAllPosts(db *gorm.DB) (*[]models.Post, error) {
	var err error
	posts := []models.Post{}
	err = db.Debug().Where(&models.Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]models.Post{}, err
	}

	return &posts, err
}

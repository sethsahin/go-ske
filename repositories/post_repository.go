package repositories

import (
	"github.com/jinzhu/gorm"
	"skeleton/models"
)

type PostRepository interface {
	GetAllPosts(db *gorm.DB) (*[]models.Post, error)
	CreatePost(db *gorm.DB, p *models.Post) (*models.Post, error)
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

func CreatePost(db *gorm.DB, p *models.Post) (*models.Post, error) {
	var err error
	err = db.Debug().Model(&models.Post{}).Create(&p).Error
	if err != nil {
		return &models.Post{}, err
	}

	return p, nil
}

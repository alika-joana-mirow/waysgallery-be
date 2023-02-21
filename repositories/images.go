package repositories

import (
	"waysGallery/models"

	"gorm.io/gorm"
)

type ImagesRepository interface {
	FindImages() ([]models.Image, error)
	GetImages(ID int) (models.Image, error)
	CreateImages(images models.Image) (models.Image, error)
	// DeleteImages(images models.Images) (models.Images, error)
}

func RepositoryImages(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindImages() ([]models.Image, error) {
	var images []models.Image
	err := r.db.Find(&images).Error

	return images, err
}

func (r *repository) GetImages(ID int) (models.Image, error) {
	var images models.Image
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.First(&images, ID).Error

	return images, err
}

func (r *repository) CreateImages(images models.Image) (models.Image, error) {
	err := r.db.Create(&images).Error

	return images, err
}

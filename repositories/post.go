package repositories

import (
	"waysGallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPost() ([]models.Post, error)
	FindImagesById(ImagesID []int) ([]models.Image, error)
	GetPost(ID int) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeletePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPost() ([]models.Post, error) {
	var post []models.Post
	err := r.db.Preload("User").Preload("Images").Find(&post).Error

	return post, err
}

func (r *repository) FindImagesById(ImagesID []int) ([]models.Image, error) {
	var images []models.Image
	err := r.db.Find(&images, ImagesID).Error

	return images, err
}

func (r *repository) GetPost(ID int) (models.Post, error) {
	var post models.Post
	err := r.db.Preload("User").Preload("Images").First(&post, ID).Error

	return post, err
}

func (r *repository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *repository) UpdatePost(post models.Post) (models.Post, error) {
	r.db.Model(&post).Association("Category").Replace(post.Images)

	err := r.db.Save(&post).Error

	return post, err
}

func (r *repository) DeletePost(post models.Post) (models.Post, error) {
	err := r.db.Delete(&post).Error

	return post, err
}

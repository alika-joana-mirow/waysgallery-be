package repositories

import (
	"waysGallery/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindProject() ([]models.Project, error)
	GetProject(ID int) (models.Project, error)
	CreateProject(project models.Project) (models.Project, error)
	FindHiredById(ID int) (models.Project, error)

	// UpdatePost(post models.Post) (models.Post, error)
	// DeletePost(post models.Post) (models.Post, error)
}

func RepositoryProject(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProject() ([]models.Project, error) {
	var project []models.Project
	err := r.db.Preload("Hired").Find(&project).Error

	return project, err
}

func (r *repository) FindHiredById(HiredID int) (models.Project, error) {
	var project models.Project
	err := r.db.Preload("Hired").Where("hired_id =?", HiredID).Find(&project).Error

	return project, err
}

func (r *repository) GetProject(ID int) (models.Project, error) {
	var project models.Project
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("Hired").First(&project, ID).Error

	return project, err
}

func (r *repository) CreateProject(project models.Project) (models.Project, error) {
	err := r.db.Create(&project).Error

	return project, err
}

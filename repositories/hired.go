package repositories

import (
	"waysGallery/models"

	"gorm.io/gorm"
)

type HiredRepository interface {
	CreateHired(hired models.Hired) (models.Hired, error)
	FindHired() ([]models.Hired, error)
	FindHiredUser(ID int) ([]models.Hired, error)
	FindHiredUserOrder(ID int) ([]models.Hired, error)
	GetHired(ID int) (models.Hired, error)
	UpdateHired(hired models.Hired) (models.Hired, error)
}

func RepositoryHired(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHired() ([]models.Hired, error) {
	var hired []models.Hired
	err := r.db.Preload("OrderBy").Preload("OrderTo").Find(&hired).Error

	return hired, err
}

func (r *repository) FindHiredUser(ID int) ([]models.Hired, error) {
	var hired []models.Hired
	err := r.db.Preload("OrderBy").Preload("OrderTo").Where("order_to_id =?", ID).Find(&hired).Error
	return hired, err
}

func (r *repository) FindHiredUserOrder(ID int) ([]models.Hired, error) {
	var hired []models.Hired
	err := r.db.Preload("OrderBy").Preload("OrderTo").Where("order_by_id =?", ID).Find(&hired).Error
	return hired, err
}

func (r *repository) GetHired(ID int) (models.Hired, error) {
	var hired models.Hired
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("OrderBy").Preload("OrderTo").First(&hired, ID).Error

	return hired, err
}

func (r *repository) CreateHired(newHired models.Hired) (models.Hired, error) {
	err := r.db.Create(&newHired).Error

	return newHired, err
}

func (r *repository) UpdateHired(hired models.Hired) (models.Hired, error) {
	err := r.db.Model(&hired).Updates(hired).Error

	return hired, err
}

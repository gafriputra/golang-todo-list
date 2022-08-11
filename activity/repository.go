package activity

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Activity, error)
	FindByID(ID int) (Activity, error)
	Save(campaign Activity) (Activity, error)
	Update(campaign Activity) (Activity, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Activity, error) {
	var activities []Activity
	err := r.db.Where("deleted_at is null").Find(&activities).Error
	return activities, err
}

func (r *repository) FindByID(ID int) (Activity, error) {
	var campaign Activity
	err := r.db.Where("deleted_at is null").Where("id = ?", ID).Find(&campaign).Error
	return campaign, err
}

func (r *repository) Save(campaign Activity) (Activity, error) {
	err := r.db.Create(&campaign).Error
	return campaign, err
}

func (r *repository) Update(campaign Activity) (Activity, error) {
	err := r.db.Save(&campaign).Error
	return campaign, err
}

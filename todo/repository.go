package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Todo, error)
	FindByID(ID int) (Todo, error)
	Save(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Where("deleted_at is null").Find(&todos).Error
	return todos, err
}

func (r *repository) FindByID(ID int) (Todo, error) {
	var todo Todo
	err := r.db.Where("deleted_at is null").Where("id = ?", ID).Find(&todo).Error
	return todo, err
}

func (r *repository) Save(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repository) Update(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

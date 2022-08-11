package todo

import (
	"errors"
	"time"
)

type Service interface {
	GetTodos() ([]Todo, error)
	GetTodoByID(input GetTodoDetailInput) (Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
	UpdateTodo(inputID GetTodoDetailInput, input CreateTodoInput) (Todo, error)
	DeleteTodo(inputID GetTodoDetailInput, input CreateTodoInput) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTodos() ([]Todo, error) {
	return s.repository.FindAll()
}

func (s *service) GetTodoByID(input GetTodoDetailInput) (Todo, error) {
	return s.repository.FindByID(input.ID)
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	now := time.Now()
	return s.repository.Save(Todo{
		ActivityGroupID: input.ActivityGroupID,
		Title:           input.Title,
		IsActive:        true,
		Priority:        "very-high",
		CreatedAt:       &now,
	})
}

func (s *service) UpdateTodo(inputID GetTodoDetailInput, input CreateTodoInput) (Todo, error) {
	activity, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return activity, err
	} else if activity.ID == 0 {
		return activity, errors.New("activity not found")
	}

	now := time.Now()
	activity.Title = input.Title
	activity.UpdatedAt = &now
	return s.repository.Update(activity)
}

func (s *service) DeleteTodo(inputID GetTodoDetailInput, input CreateTodoInput) (Todo, error) {
	activity, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return activity, err
	} else if activity.ID == 0 {
		return activity, errors.New("activity not found")
	}

	now := time.Now()
	activity.DeletedAt = &now
	return s.repository.Update(activity)
}

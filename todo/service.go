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
	DeleteTodo(ID int) (Todo, error)
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
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	} else if todo.ID == 0 {
		return todo, errors.New("todo not found")
	}

	now := time.Now()
	todo.Title = input.Title
	todo.UpdatedAt = &now
	return s.repository.Update(todo)
}

func (s *service) DeleteTodo(ID int) (Todo, error) {
	todo, err := s.repository.FindByID(ID)
	if err != nil {
		return todo, err
	} else if todo.ID == 0 {
		return todo, errors.New("todo not found")
	}

	now := time.Now()
	todo.DeletedAt = &now
	return s.repository.Update(todo)
}

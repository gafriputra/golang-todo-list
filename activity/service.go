package activity

import (
	"errors"
	"time"
)

type Service interface {
	GetActivities() ([]Activity, error)
	GetActivityByID(input GetActivityDetailInput) (Activity, error)
	CreateActivity(input CreateActivityInput) (Activity, error)
	UpdateActivity(inputID GetActivityDetailInput, input CreateActivityInput) (Activity, error)
	DeleteActivity(id int) (Activity, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetActivities() ([]Activity, error) {
	return s.repository.FindAll()
}

func (s *service) GetActivityByID(input GetActivityDetailInput) (Activity, error) {
	return s.repository.FindByID(input.ID)
}

func (s *service) CreateActivity(input CreateActivityInput) (Activity, error) {
	now := time.Now()
	return s.repository.Save(Activity{
		Email:     input.Email,
		Title:     input.Title,
		CreatedAt: &now,
	})
}

func (s *service) UpdateActivity(inputID GetActivityDetailInput, input CreateActivityInput) (Activity, error) {
	activity, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return activity, err
	} else if activity.ID == 0 {
		return activity, errors.New("activity not found")
	}

	now := time.Now()
	activity.Email = input.Email
	activity.Title = input.Title
	activity.UpdatedAt = &now
	return s.repository.Update(activity)
}

func (s *service) DeleteActivity(id int) (Activity, error) {
	activity, err := s.repository.FindByID(id)
	if err != nil {
		return activity, err
	} else if activity.ID == 0 {
		return activity, errors.New("activity not found")
	}

	now := time.Now()
	activity.DeletedAt = &now
	return s.repository.Update(activity)
}

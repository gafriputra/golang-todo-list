package todo

import (
	"time"
)

type Todo struct {
	ID              int
	ActivityGroupID int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}

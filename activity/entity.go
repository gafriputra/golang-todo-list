package activity

import (
	"time"
)

type Activity struct {
	ID        int
	Email     string
	Title     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

package domain

import (
	"time"
)

type Task struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Completed   bool
}

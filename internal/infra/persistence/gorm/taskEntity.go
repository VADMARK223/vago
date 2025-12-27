package gorm

import "time"

type TaskEntity struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Completed   bool      `gorm:"default:false"`

	UserID int64
}

func (TaskEntity) TableName() string {
	return "tasks"
}

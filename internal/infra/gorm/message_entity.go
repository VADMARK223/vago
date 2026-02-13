package gorm

import (
	"time"
)

type MessageEntity struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	Type      string    `gorm:"column:message_type"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (MessageEntity) TableName() string {
	return "messages"
}

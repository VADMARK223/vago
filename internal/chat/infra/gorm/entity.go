package gorm

import "time"

type MessageEntity struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (MessageEntity) TableName() string {
	return "messages"
}

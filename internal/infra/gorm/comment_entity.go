package gorm

import "time"

type CommentEntity struct {
	ID         int64     `gorm:"primaryKey;column:id"`
	QuestionID int64     `gorm:"column:question_id;not null;index"`
	ParentID   *int64    `gorm:"column:parent_id;index"`
	AuthorID   int64     `gorm:"column:author_id;not null"`
	Content    string    `gorm:"column:content;type:text;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (e *CommentEntity) TableName() string { return "comments" }

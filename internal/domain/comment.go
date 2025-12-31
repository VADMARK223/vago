package domain

import "time"

type Comment struct {
	ID             int64
	QuestionID     int64
	ParentID       *int64 // Если nil, то комментарий к вопросу, иначе к комментарию
	AuthorID       int64
	AuthorUsername string
	Content        string
	CreatedAt      time.Time
	Children       []*Comment // заполняется в сервисе
}

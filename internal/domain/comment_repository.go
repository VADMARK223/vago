package domain

import (
	"context"
)

type CommentRepo interface {
	List(ctx context.Context) ([]*Comment, error)
	ListByQuestionID(ctx context.Context, questionID int64) ([]*Comment, error)

	Create(ctx context.Context, c *Comment) (*Comment, error)
	GetByID(ctx context.Context, id int64) (*Comment, error)
}

package repository

import (
	"context"
	"vago/internal/domain"
)

type CommentRepo interface {
	List(ctx context.Context) ([]*domain.Comment, error)
	ListByQuestionID(ctx context.Context, questionID int64) ([]*domain.Comment, error)
}

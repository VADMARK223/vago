package repository

import (
	"context"
	"vago/internal/domain"
)

type CommentRepo interface {
	List(ctx context.Context) ([]*domain.Comment, error)
}

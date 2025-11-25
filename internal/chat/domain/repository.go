package domain

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, dto MessageDTO) error
	ListAll(ctx context.Context) ([]*Message, error)
}

package domain

import (
	"context"
)

type MessageRepository interface {
	Save(ctx context.Context, message *Message) (uint, error)
	ListAll(ctx context.Context) ([]*Message, error)
	DeleteMessage(id uint) error
	DeleteAll() error
}

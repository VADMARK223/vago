package domain

import (
	"context"
)

type MessageRepository interface {
	Save(ctx context.Context, message Message) (MessageID, error)
	ListAll(ctx context.Context) ([]Message, error)
	DeleteMessage(id int64) error
	DeleteAll() error
}

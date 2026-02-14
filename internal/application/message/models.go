package message

import (
	"time"
	"vago/internal/domain"
)

type WithUsername struct {
	ID          domain.MessageID
	AuthorID    domain.UserID
	Username    string
	Body        string
	SentAt      time.Time
	MessageType string
}

package chat

import "vago/internal/domain"

type MessageCreateDTO struct {
	AuthorID    domain.UserID `json:"author_id"`
	Body        string        `json:"body"`
	MessageType string        `json:"type"`
}

type MessageDTO struct {
	ID          domain.MessageID `json:"id"`
	AuthorID    domain.UserID    `json:"author_id"`
	Username    string           `json:"username"`
	Body        string           `json:"body"`
	SentAt      string           `json:"sent_at"`
	MessageType string           `json:"type"`
}

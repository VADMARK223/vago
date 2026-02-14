package chat

import "vago/internal/domain"

type MessageCreateDTO struct {
	AuthorID    domain.UserID `json:"author_id"`
	Body        string        `json:"body"`
	MessageType string        `json:"type"`
}

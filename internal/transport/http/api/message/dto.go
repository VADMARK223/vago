package message

import (
	"time"
	"vago/internal/application/message"
	"vago/internal/domain"
)

type MessageApiDTO struct {
	ID          domain.MessageID `json:"id"`
	AuthorID    domain.UserID    `json:"authorId"`
	Username    string           `json:"username"`
	Body        string           `json:"body"`
	SentAt      string           `json:"sentAt"`
	MessageType string           `json:"type"`
}

type MessagesApiDTO struct {
	Messages []MessageApiDTO `json:"messages"`
}

func messageToDTO(m message.WithUsername) MessageApiDTO {
	return MessageApiDTO{
		ID:          m.ID,
		AuthorID:    m.AuthorID,
		Username:    m.Username,
		MessageType: m.MessageType,
		Body:        m.Body,
		SentAt:      m.SentAt.UTC().Format(time.RFC3339Nano),
	}
}

func MessagesToDTO(messages []message.WithUsername) []MessageApiDTO {
	result := make([]MessageApiDTO, 0, len(messages))
	for _, m := range messages {
		result = append(result, messageToDTO(m))
	}
	return result
}

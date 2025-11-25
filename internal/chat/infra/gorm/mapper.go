package gorm

import (
	"vago/internal/chat/domain"
)

func messageToEnty(dto domain.MessageDTO) *MessageEntity {
	return &MessageEntity{
		UserID:    uint(dto.Author),
		Content:   string(dto.Body),
		CreatedAt: dto.SentAt,
	}
}

func messageToDomain(e *MessageEntity) (*domain.Message, error) {
	body, err := domain.NewBody(e.Content)
	if err != nil {
		return nil, err
	}

	return domain.New(
		e.ID,
		domain.UserID(e.UserID),
		body,
		e.CreatedAt,
	)
}

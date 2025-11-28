package gorm

import (
	"vago/internal/chat/domain"
	"vago/pkg/timex"
)

func messageToEnty(dto domain.MessageDTO) *MessageEntity {
	t, err := timex.Parse(dto.SentAt)
	if err != nil {
		panic(err)
	}
	return &MessageEntity{
		UserID:    uint(dto.AuthorID),
		Content:   string(dto.Body),
		CreatedAt: t,
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

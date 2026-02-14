package chat

import (
	"context"

	"vago/internal/domain"
	"vago/internal/transport/http/api/message"
	"vago/pkg/timex"
)

type Service struct {
	msgRepo  domain.MessageRepository
	userRepo domain.UserRepository
}

func NewService(messageRepo domain.MessageRepository, userRepo domain.UserRepository) *Service {
	return &Service{
		msgRepo:  messageRepo,
		userRepo: userRepo,
	}
}

func (s *Service) CreateMessage(ctx context.Context, dto MessageCreateDTO) (message.MessageApiDTO, error) {
	msg := domain.NewMessage(dto.AuthorID, domain.Body(dto.Body), dto.MessageType)
	id, err := s.msgRepo.Save(ctx, msg)
	if err != nil {
		return message.MessageApiDTO{}, err
	}

	user, err := s.userRepo.GetByID(dto.AuthorID)
	if err != nil {
		return message.MessageApiDTO{}, err
	}

	return message.MessageApiDTO{
		ID:          id,
		AuthorID:    domain.UserID(user.ID),
		Username:    user.Username,
		Body:        dto.Body,
		MessageType: dto.MessageType,
		SentAt:      timex.Format(msg.SentAt()),
	}, nil
}

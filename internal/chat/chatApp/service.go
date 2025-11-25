package chatApp

import (
	"context"
	"time"
	"vago/internal/chat/domain"
)

type Service struct {
	repo domain.Repository
}

func NewMessageSvc(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SendMessage(ctx context.Context, author domain.UserID, body string) error {
	b, errBody := domain.NewBody(body)
	if errBody != nil {
		return errBody
	}

	dto := domain.MessageDTO{
		Author: author,
		Body:   b,
		SentAt: time.Now(),
	}

	if err := s.repo.Save(ctx, dto); err != nil {
		return err
	}

	return nil
}

func (s *Service) LastMessages(ctx context.Context) ([]*domain.Message, error) {
	return s.repo.ListAll(ctx)
}

func (s *Service) DeleteMessage(id uint) error {
	return s.repo.DeleteMessage(id)
}

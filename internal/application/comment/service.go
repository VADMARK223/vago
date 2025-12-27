package comment

import (
	"context"
	"vago/internal/domain"
	"vago/internal/domain/repository"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	repo repository.CommentRepo
}

func NewService(repo *gorm.CommentRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) All() ([]*domain.Comment, error) {
	ctx := context.TODO()
	comments, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

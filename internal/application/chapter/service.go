package chapter

import (
	"vago/internal/domain"
	"vago/internal/domain/repository"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	repo repository.ChapterRepository
}

func NewService(repo *gorm.ChapterRepo) *Service {
	return &Service{repo: *repo}
}

func (s *Service) All() ([]*domain.Chapter, error) {
	chapters, err := s.repo.All()

	if err != nil {
		return nil, err
	}

	return chapters, nil
}

package chapter

import (
	"vago/internal/domain"
	"vago/internal/infra/gorm"
)

type Service struct {
	repo domain.ChapterRepository
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

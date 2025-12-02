package topic

import (
	"vago/internal/domain"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	repo domain.TopicRepository
}

func NewService(repo *gorm.TopicRepo) *Service {
	return &Service{repo: *repo}
}

func (s *Service) All() ([]*domain.Topic, error) {
	questions, err := s.repo.All()

	if err != nil {
		return nil, err
	}

	return questions, nil
}

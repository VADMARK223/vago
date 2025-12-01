package quiz

import (
	"vago/internal/domain"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	repo gorm.QuestionRepo
}

func NewService(repo *gorm.QuestionRepo) *Service {
	return &Service{repo: *repo}
}

func (s *Service) AllQuestions() ([]*domain.Question, error) {
	questions, err := s.repo.All()

	if err != nil {
		return nil, err
	}

	return questions, nil
}

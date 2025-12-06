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

func (s *Service) DeleteAll() error {
	return s.repo.DeleteAll()
}

func (s *Service) RandomQuestion(id *uint) (*domain.Question, error) {
	if id == nil {
		return s.repo.Random()
	}

	return s.repo.GetByID(*id)
}

func (s *Service) ToPublic(q *domain.Question) QuestionPublic {
	res := QuestionPublic{
		ID:          q.ID,
		Text:        q.Text,
		Code:        q.Code,
		Explanation: q.Explanation,
	}

	for _, a := range q.Answers {
		res.Answers = append(res.Answers, AnswerPublic{
			ID:   a.ID,
			Text: a.Text,
		})
	}
	return res
}

func (s *Service) CheckAnswer(qID, aID uint) bool {
	q, _ := s.repo.GetByID(qID)

	for _, a := range q.Answers {
		if a.ID == aID {
			return a.IsCorrect
		}
	}
	return false
}

package quiz

import (
	"vago/internal/domain"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	questionRepo gorm.QuestionRepo
	topicRepo    gorm.TopicRepo
}

func NewService(questionRepo *gorm.QuestionRepo, topicRepo *gorm.TopicRepo) *Service {
	return &Service{questionRepo: *questionRepo, topicRepo: *topicRepo}
}

func (s *Service) AllQuestions() ([]*domain.Question, error) {
	questions, err := s.questionRepo.All()

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *Service) DeleteAll() error {
	return s.questionRepo.DeleteAll()
}

func (s *Service) RandomPublicQuestion(id *uint) QuestionPublic {
	question, err := s.randomQuestion(id)
	if err != nil {
		panic(err)
	}
	topic, err := s.topicRepo.GetByID(question.TopicID)
	if err != nil {
		panic(err)
	}

	return s.toPublic(question, topic.Name)
}

func (s *Service) randomQuestion(id *uint) (*domain.Question, error) {
	if id == nil {
		return s.questionRepo.Random()
	}

	return s.questionRepo.GetByID(*id)
}

func (s *Service) toPublic(q *domain.Question, topicName string) QuestionPublic {
	if q == nil {
		return QuestionPublic{
			ID:   uint(0),
			Text: "Ошибка поиска вопроса",
		}
	}
	res := QuestionPublic{
		ID:          q.ID,
		Text:        q.Text,
		Code:        q.Code,
		Explanation: q.Explanation,
		TopicName:   topicName,
	}

	for _, a := range q.Answers {
		res.Answers = append(res.Answers, AnswerPublic{
			ID:   a.ID,
			Text: a.Text,
		})
	}
	return res
}

func (s *Service) CheckAnswer(qID, aID uint) (bool, string) {
	q, _ := s.questionRepo.GetByID(qID)

	for _, a := range q.Answers {
		if a.ID == aID {
			return a.IsCorrect, q.Explanation
		}
	}
	return false, q.Explanation
}

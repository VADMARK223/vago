package test

import (
	"vago/internal/domain"
	gorm2 "vago/internal/infra/gorm"
)

type Service struct {
	questionRepo gorm2.QuestionRepo
	topicRepo    gorm2.TopicRepo
}

func NewService(questionRepo *gorm2.QuestionRepo, topicRepo *gorm2.TopicRepo) *Service {
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

func (s *Service) RandomID() (int64, error) {
	return s.questionRepo.RandomID()
}

// PublicQuestion если id пустой, возвращаем случайный вопрос, иначе получаем вопрос по его id
func (s *Service) PublicQuestion(id *int64) QuestionPublic {
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

func (s *Service) randomQuestion(id *int64) (*domain.Question, error) {
	if id == nil {
		return s.questionRepo.Random()
	}

	return s.questionRepo.GetByID(*id)
}

func (s *Service) toPublic(q *domain.Question, topicName string) QuestionPublic {
	if q == nil {
		return QuestionPublic{
			ID:   0,
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

func (s *Service) CheckAnswer(qID, aID int64) (bool, string, error) {
	q, err := s.questionRepo.GetByID(qID)

	if err != nil {
		return false, "", err
	}

	for _, a := range q.Answers {
		if a.ID == aID {
			return a.IsCorrect, q.Explanation, nil
		}
	}
	return false, q.Explanation, err
}

func (s *Service) GetQuestionsByTopic(topicId int64) ([]*domain.Question, error) {
	return s.questionRepo.FindByTopicID(topicId)
}

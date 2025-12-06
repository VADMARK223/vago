package gorm

import (
	"math/rand"
	"time"
	"vago/internal/domain"

	"gorm.io/gorm"
)

type QuestionRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

func (q QuestionRepo) DeleteAll() error {
	return q.db.Exec("TRUNCATE TABLE questions RESTART IDENTITY CASCADE").Error
}

func (q QuestionRepo) All() ([]*domain.Question, error) {
	var entities []QuestionEntity

	err := q.db.Preload("Answers").Find(&entities).Error

	if err != nil {
		return nil, err
	}
	result := make([]*domain.Question, 0, len(entities))

	for _, e := range entities {
		q := &domain.Question{
			ID:   e.ID,
			Text: e.Text,
		}

		for _, a := range e.Answers {
			q.Answers = append(q.Answers, domain.Answer{
				ID:        a.ID,
				Text:      a.Text,
				IsCorrect: a.IsCorrect,
			})
		}

		result = append(result, q)
	}

	return result, nil
}

func (q QuestionRepo) Random() (*domain.Question, error) {
	var entity QuestionEntity

	err := q.db.
		Preload("Answers").
		Order("RANDOM()").
		Limit(1).
		First(&entity).Error

	if err != nil {
		return nil, err
	}

	// Маппинг в доменную модель
	res := &domain.Question{
		ID:   entity.ID,
		Text: entity.Text,
	}

	for _, a := range entity.Answers {
		res.Answers = append(res.Answers, domain.Answer{
			ID:        a.ID,
			Text:      a.Text,
			IsCorrect: a.IsCorrect,
		})
	}

	shuffleAnswers(res.Answers)

	return res, nil
}

func shuffleAnswers(a []domain.Answer) {
	rand.Seed(time.Now().UnixNano()) // однократная инициализация
	n := len(a)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func (q QuestionRepo) GetByID(id uint) (*domain.Question, error) {
	var entity QuestionEntity

	err := q.db.Preload("Answers").
		First(&entity, id).Error

	if err != nil {
		return nil, err
	}

	result := &domain.Question{
		ID:   entity.ID,
		Text: entity.Text,
	}

	for _, a := range entity.Answers {
		result.Answers = append(result.Answers, domain.Answer{
			ID:        a.ID,
			Text:      a.Text,
			IsCorrect: a.IsCorrect,
		})
	}

	return result, nil
}

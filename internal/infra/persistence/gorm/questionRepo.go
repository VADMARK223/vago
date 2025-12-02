package gorm

import (
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

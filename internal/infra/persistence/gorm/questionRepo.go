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

func (q QuestionRepo) All() ([]*domain.Question, error) {
	var entities []QuestionEntity

	err := q.db.Find(&entities).Error

	if err != nil {
		return nil, err
	}
	result := make([]*domain.Question, 0, len(entities))

	for _, entity := range entities {
		q := &domain.Question{
			ID:   entity.ID,
			Text: entity.Text,
		}
		result = append(result, q)
	}

	return result, nil
}

package gorm

import (
	"vago/internal/domain"

	"gorm.io/gorm"
)

type TopicRepo struct {
	db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) *TopicRepo {
	return &TopicRepo{db: db}
}

func (r TopicRepo) All() ([]*domain.Topic, error) {
	var entities []TopicEntity

	err := r.db.Find(&entities).Error

	if err != nil {
		return nil, err
	}
	result := make([]*domain.Topic, 0, len(entities))

	for _, e := range entities {
		q := &domain.Topic{
			ID:   e.ID,
			Name: e.Name,
		}

		result = append(result, q)
	}

	return result, nil
}

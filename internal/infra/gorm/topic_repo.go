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

func (r *TopicRepo) All() ([]*domain.Topic, error) {
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

func (r *TopicRepo) GetByID(id int64) (*domain.Topic, error) {
	var entity TopicEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &domain.Topic{
		ID:   entity.ID,
		Name: entity.Name,
	}, nil
}

func (r *TopicRepo) AllWithCount() ([]domain.TopicWithCount, error) {
	var topics []domain.TopicWithCount

	err := r.db.
		Table("topics t").
		Select("t.id, t.name, COUNT(q.id) AS questions_count").
		Joins("LEFT JOIN questions q ON q.topic_id = t.id").
		Group("t.id, t.name").
		Order("t.id").
		Scan(&topics).Error

	return topics, err
}

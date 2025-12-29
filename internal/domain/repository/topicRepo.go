package repository

import "vago/internal/domain"

type TopicRepository interface {
	All() ([]*domain.Topic, error)
	AllWithCount() ([]domain.TopicWithCount, error)
	GetByID(id int64) (*domain.Topic, error)
}

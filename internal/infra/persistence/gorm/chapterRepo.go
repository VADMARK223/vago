package gorm

import (
	"vago/internal/domain"

	"gorm.io/gorm"
)

type ChapterRepo struct {
	db *gorm.DB
}

func NewChapterRepo(db *gorm.DB) *ChapterRepo {
	return &ChapterRepo{db: db}
}

func (r ChapterRepo) All() ([]*domain.Chapter, error) {
	var entities []ChapterEntity

	err := r.db.Find(&entities).Error

	if err != nil {
		return nil, err
	}
	result := make([]*domain.Chapter, 0, len(entities))

	for _, e := range entities {
		q := &domain.Chapter{
			ID:   e.ID,
			Name: e.Name,
		}

		result = append(result, q)
	}

	return result, nil
}

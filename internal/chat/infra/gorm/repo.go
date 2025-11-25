package gorm

import (
	"context"
	"vago/internal/chat/domain"

	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Save(ctx context.Context, dto domain.MessageDTO) error {
	entity := messageToEnty(dto)

	err := r.db.WithContext(ctx).Create(entity).Error

	//domain, errDomain := messageToDomain(entity)

	return err
}

func (r *MessageRepo) ListAll(ctx context.Context) ([]*domain.Message, error) {
	var entities []MessageEntity

	err := r.db.WithContext(ctx).Find(&entities).Error

	/*err := r.db.WithContext(ctx).
	Order("id DESC").
	Limit(limit).
	Find(&entities).Error*/

	if err != nil {
		return nil, err
	}

	result := make([]*domain.Message, 0, len(entities))
	for _, entity := range entities {
		m, errMapping := messageToDomain(&entity)
		if errMapping != nil {
			return nil, errMapping
		}
		result = append(result, m)
	}

	return result, nil
}

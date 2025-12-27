package gorm

import (
	"context"
	"fmt"
	"time"
	"vago/internal/domain"

	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Save(ctx context.Context, m *domain.Message) (int64, error) {
	entity := MessageEntity{
		UserID:    int64(m.Author()),
		Content:   string(m.Body()),
		Type:      m.MessageType,
		CreatedAt: time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return 0, err
	}

	return entity.ID, nil
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
		m := domain.NewMessage(domain.UserID(entity.UserID), domain.Body(entity.Content), entity.Type)
		m.ID = entity.ID
		m.MessageType = entity.Type
		m.SetSentAt(entity.CreatedAt)
		result = append(result, m)
	}

	return result, nil
}

func (r *MessageRepo) DeleteMessage(id int64) error {
	if err := r.db.Delete(&MessageEntity{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	return nil
}

func (r *MessageRepo) DeleteAll() error {
	return r.db.Exec("TRUNCATE TABLE messages RESTART IDENTITY").Error
}

package gorm

import (
	"context"
	"vago/internal/domain"

	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (repo *CommentRepo) List(ctx context.Context) ([]*domain.Comment, error) {
	var ents []*CommentEntity
	err := repo.db.WithContext(ctx).Find(&ents).Error

	if err != nil {
		return nil, err
	}

	out := make([]*domain.Comment, 0, len(ents))
	for _, e := range ents {
		out = append(out, &domain.Comment{
			ID:         e.ID,
			QuestionID: e.QuestionID,
			ParentID:   e.ParentID,
			AuthorID:   e.AuthorID,
			Content:    e.Content,
			CreatedAt:  e.CreatedAt,
		})
	}
	return out, nil
}

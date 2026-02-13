package gorm

import (
	"context"
	"time"
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

type commentRow struct {
	ID             int64
	QuestionID     int64
	ParentID       *int64
	AuthorID       int64
	AuthorUsername string
	Content        string
	CreatedAt      time.Time
}

func (r *CommentRepo) ListByQuestionID(ctx context.Context, qid int64) ([]*domain.Comment, error) {
	var rows []commentRow
	err := r.db.WithContext(ctx).
		Table("comments c").
		Select(`
			c.id,
			c.question_id,
			c.parent_id,
			c.author_id,
			u.username AS author_username,
			c.content,
			c.created_at
		`).
		Joins("JOIN users u ON u.id = c.author_id").
		Where("question_id = ?", qid).
		Order("created_at ASC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}

	out := make([]*domain.Comment, 0, len(rows))
	for _, e := range rows {
		out = append(out, &domain.Comment{
			ID:             e.ID,
			QuestionID:     e.QuestionID,
			ParentID:       e.ParentID,
			AuthorID:       e.AuthorID,
			AuthorUsername: e.AuthorUsername,
			Content:        e.Content,
			CreatedAt:      e.CreatedAt,
		})
	}
	return out, nil
}

func (r *CommentRepo) List(ctx context.Context) ([]*domain.Comment, error) {
	var ents []*CommentEntity
	err := r.db.WithContext(ctx).Find(&ents).Error

	if err != nil {
		return nil, err
	}

	out := make([]*domain.Comment, 0, len(ents))
	for _, e := range ents {
		out = append(out, commentEntityToDomain(e))
	}
	return out, nil
}

func (r *CommentRepo) Create(ctx context.Context, c *domain.Comment) (*domain.Comment, error) {
	entity := CommentEntity{
		QuestionID: c.QuestionID,
		ParentID:   c.ParentID,
		AuthorID:   c.AuthorID,
		Content:    c.Content,
	}

	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return nil, err
	}

	c.ID = entity.ID
	c.CreatedAt = entity.CreatedAt
	return c, nil
}

func (r *CommentRepo) GetByID(ctx context.Context, id int64) (*domain.Comment, error) {
	var e CommentEntity
	if err := r.db.WithContext(ctx).First(&e, id).Error; err != nil {
		return nil, err
	}

	return commentEntityToDomain(&e), nil
}

func commentEntityToDomain(e *CommentEntity) *domain.Comment {
	return &domain.Comment{
		ID:         e.ID,
		QuestionID: e.QuestionID,
		ParentID:   e.ParentID,
		AuthorID:   e.AuthorID,
		Content:    e.Content,
		CreatedAt:  e.CreatedAt,
	}
}

package comment

import (
	"context"
	"vago/internal/domain"
	"vago/internal/domain/repository"
	"vago/internal/infra/persistence/gorm"
)

type Service struct {
	repo repository.CommentRepo
}

func NewService(repo *gorm.CommentRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) All() ([]*domain.Comment, error) {
	ctx := context.TODO()
	comments, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *Service) ListByQuestionID(questionID int64) ([]*domain.Comment, int, error) {
	ctx := context.TODO()
	comments, err := s.repo.ListByQuestionID(ctx, questionID)
	if err != nil {
		return nil, 0, err
	}

	result := buildTree(comments)

	return result, len(comments), nil
}

func buildTree(comments []*domain.Comment) []*domain.Comment {
	byID := make(map[int64]*domain.Comment, len(comments))
	roots := make([]*domain.Comment, 0)

	for _, c := range comments {
		c.Children = nil
		byID[c.ID] = c
	}

	for _, c := range comments {
		if c.ParentID == nil {
			roots = append(roots, c)
			continue
		}

		parent := byID[*c.ParentID]
		if parent != nil {
			parent.Children = append(parent.Children, c)
		}
	}

	return roots
}

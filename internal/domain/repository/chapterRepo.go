package repository

import "vago/internal/domain"

type ChapterRepository interface {
	All() ([]*domain.Chapter, error)
}

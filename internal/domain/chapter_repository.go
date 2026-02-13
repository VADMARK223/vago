package domain

type ChapterRepository interface {
	All() ([]*Chapter, error)
}

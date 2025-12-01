package domain

type QuestionRepository interface {
	All() ([]*Question, error)
}

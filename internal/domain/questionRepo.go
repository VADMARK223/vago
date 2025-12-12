package domain

type QuestionRepository interface {
	All() ([]*Question, error)
	DeleteAll() error
	Random() (*Question, error)
	GetByID(id uint) (*Question, error)
	RandomID() (uint, error)
	FindByTopicID(topicID uint) ([]*Question, error)
}

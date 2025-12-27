package domain

type QuestionRepository interface {
	All() ([]*Question, error)
	DeleteAll() error
	Random() (*Question, error)
	GetByID(id int64) (*Question, error)
	RandomID() (int64, error)
	FindByTopicID(topicID uint) ([]*Question, error)
}

package domain

type TopicRepository interface {
	All() ([]*Topic, error)
	AllWithCount() ([]TopicWithCount, error)
	GetByID(id int64) (*Topic, error)
}

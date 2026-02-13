package domain

type TopicRepository interface {
	All() ([]*Topic, error)
	GetByID(id int64) (*Topic, error)
	AllWithCount() ([]TopicWithCount, error)
}

package domain

type TopicRepository interface {
	All() ([]*Topic, error)
	AllWithCount() ([]TopicWithCount, error)
}

package domain

type TopicRepository interface {
	All() ([]*Topic, error)
}

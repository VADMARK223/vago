package domain

type Topic struct {
	ID   int64
	Name string
}

type TopicWithCount struct {
	ID             int64
	Name           string
	QuestionsCount int64
}

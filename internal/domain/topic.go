package domain

type Topic struct {
	ID   uint
	Name string
}

type TopicWithCount struct {
	ID             uint
	Name           string
	QuestionsCount uint
}

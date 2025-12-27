package quiz

type QuestionPublic struct {
	ID          int64
	Text        string
	Code        string
	Explanation string
	TopicName   string
	Answers     []AnswerPublic
}

type AnswerPublic struct {
	ID   int64
	Text string
}

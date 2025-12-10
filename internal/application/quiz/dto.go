package quiz

type QuestionPublic struct {
	ID          uint
	Text        string
	Code        string
	Explanation string
	TopicName   string
	Answers     []AnswerPublic
}

type AnswerPublic struct {
	ID   uint
	Text string
}

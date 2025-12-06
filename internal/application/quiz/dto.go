package quiz

type QuestionPublic struct {
	ID          uint
	Text        string
	Code        string
	Explanation string
	Answers     []AnswerPublic
}

type AnswerPublic struct {
	ID   uint
	Text string
}

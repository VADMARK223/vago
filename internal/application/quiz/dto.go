package quiz

type QuestionPublic struct {
	ID      uint
	Text    string
	Answers []AnswerPublic
}

type AnswerPublic struct {
	ID   uint
	Text string
}

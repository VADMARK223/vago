package domain

type Question struct {
	ID          uint
	TopicID     uint
	Text        string
	Code        string
	Explanation string
	Answers     []Answer
}

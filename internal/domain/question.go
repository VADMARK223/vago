package domain

type Question struct {
	ID          int64
	TopicID     int64
	Text        string
	Code        string
	Explanation string
	Answers     []Answer
}

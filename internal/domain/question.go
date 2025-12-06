package domain

type Question struct {
	ID          uint
	Text        string
	Code        string
	Explanation string
	Answers     []Answer
}

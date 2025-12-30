package comment

type CreateCommentDTO struct {
	QuestionID int64
	ParentID   *int64
	AuthorID   int64
	Content    string
}

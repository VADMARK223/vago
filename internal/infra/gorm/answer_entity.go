package gorm

type AnswerEntity struct {
	ID         int64  `gorm:"primaryKey"`
	QuestionID int64  `gorm:"column:question_id"`
	Text       string `gorm:"column:text"`
	IsCorrect  bool   `gorm:"column:is_correct"`
}

func (AnswerEntity) TableName() string {
	return "answers"
}

package gorm

type AnswerEntity struct {
	ID         uint   `gorm:"primaryKey"`
	QuestionID uint   `gorm:"column:question_id"`
	Text       string `gorm:"column:text"`
	IsCorrect  bool   `gorm:"column:is_correct"`
}

func (AnswerEntity) TableName() string {
	return "answers"
}

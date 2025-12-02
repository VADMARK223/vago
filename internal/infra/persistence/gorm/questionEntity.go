package gorm

type QuestionEntity struct {
	ID      uint           `gorm:"primaryKey"`
	Text    string         `gorm:"column:text"`
	Answers []AnswerEntity `gorm:"foreignKey:QuestionID"`
}

func (QuestionEntity) TableName() string {
	return "questions"
}

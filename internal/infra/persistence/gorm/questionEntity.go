package gorm

type QuestionEntity struct {
	ID          uint           `gorm:"primaryKey"`
	Text        string         `gorm:"column:text"`
	Code        string         `gorm:"column:code"`
	Explanation string         `gorm:"column:explanation"`
	Answers     []AnswerEntity `gorm:"foreignKey:QuestionID"`
}

func (QuestionEntity) TableName() string {
	return "questions"
}

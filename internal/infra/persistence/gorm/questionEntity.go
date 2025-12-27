package gorm

type QuestionEntity struct {
	ID          int64          `gorm:"primaryKey"`
	TopicID     int64          `gorm:"column:topic_id"`
	Text        string         `gorm:"column:text"`
	Code        string         `gorm:"column:code"`
	Explanation string         `gorm:"column:explanation"`
	Answers     []AnswerEntity `gorm:"foreignKey:QuestionID"`
}

func (QuestionEntity) TableName() string {
	return "questions"
}

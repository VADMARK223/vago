package gorm

type QuestionEntity struct {
	ID   uint   `gorm:"primaryKey"`
	Text string `gorm:"column:text"`
}

func (QuestionEntity) TableName() string {
	return "questions"
}

package gorm

type TopicEntity struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func (TopicEntity) TableName() string {
	return "topics"
}

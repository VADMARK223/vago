package gorm

type TopicEntity struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}

func (TopicEntity) TableName() string {
	return "topics"
}

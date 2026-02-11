package gorm

type ChapterEntity struct {
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Order int64  `gorm:"column:order"`
}

func (ChapterEntity) TableName() string {
	return "chapters"
}

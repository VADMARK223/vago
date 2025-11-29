package gorm

import (
	"vago/internal/domain"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) domain.TaskRepository {
	return &TaskRepo{db: db}
}

func (r TaskRepo) GetAll() ([]domain.Task, error) {
	var entities []TaskEntity
	err := r.db.Find(&entities).Error

	result := make([]domain.Task, 0, len(entities))
	for _, entity := range entities {
		result = append(result, domain.Task{
			ID:   entity.ID,
			Name: entity.Name,
		})
	}

	return result, err
}

func (r TaskRepo) GetAllByUserID(ID uint) ([]domain.Task, error) {
	var entities []TaskEntity
	err := r.db.Where("user_id = ?", ID).Find(&entities).Error

	result := make([]domain.Task, 0, len(entities))
	for _, entity := range entities {
		result = append(result, domain.Task{
			ID:          entity.ID,
			Name:        entity.Name,
			Description: entity.Description,
			Completed:   entity.Completed,
			CreatedAt:   entity.CreatedAt,
		})
	}

	return result, err
}

func (r TaskRepo) UpdateCompleted(taskID, userID uint, completed bool) error {
	taskEntity := TaskEntity{}
	if err := r.db.Where("id = ? AND user_id = ?", taskID, userID).First(&taskEntity).Error; err != nil {
		return err
	}

	taskEntity.Completed = completed
	return r.db.Save(&taskEntity).Error
}

package gorm

import (
	"errors"
	"vago/internal/domain"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) domain.TaskRepository {
	return &TaskRepo{db: db}
}

func (r TaskRepo) PostTask(name string, desc string, completed bool, userID int64) error {
	t := TaskEntity{
		Name:        name,
		Description: desc,
		Completed:   completed,
		UserID:      userID,
	}

	if err := r.db.Create(&t).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (r TaskRepo) GetAllByUserID(ID int64) ([]domain.Task, error) {
	var entities []TaskEntity
	err := r.db.
		Where("user_id = ?", ID).
		Order("completed ASC").
		Order("created_at DESC").
		Find(&entities).Error

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

func (r TaskRepo) UpdateCompleted(taskID, userID int64, completed bool) error {
	taskEntity := TaskEntity{}
	if err := r.db.Where("id = ? AND user_id = ?", taskID, userID).First(&taskEntity).Error; err != nil {
		return err
	}

	taskEntity.Completed = completed
	return r.db.Save(&taskEntity).Error
}

func (r TaskRepo) DeleteTask(taskID int64) error {
	if err := r.db.Delete(&domain.Task{}, taskID).Error; err != nil {
		return errors.New(err.Error())
	}

	return nil
}

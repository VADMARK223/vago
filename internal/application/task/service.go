package task

import (
	"vago/internal/domain"
)

type Service struct {
	repo domain.TaskRepository
}

func NewService(repo domain.TaskRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks() ([]domain.Task, error) {
	return s.repo.GetAll()
}

func (s *Service) GetAllByUser(userID uint) ([]domain.Task, error) {
	return s.repo.GetAllByUserID(userID)
}

func (s *Service) UpdateCompleted(taskID, userID uint, completed bool) error {
	return s.repo.UpdateCompleted(taskID, userID, completed)
}

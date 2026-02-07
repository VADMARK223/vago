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

func (s *Service) PostTask(name string, desc string, completed bool, userID int64) error {
	return s.repo.PostTask(name, desc, completed, userID)
}

func (s *Service) GetAllByUser(userID int64) ([]domain.Task, error) {
	return s.repo.GetAllByUserID(userID)
}

func (s *Service) UpdateCompleted(taskID, userID int64, completed bool) error {
	return s.repo.UpdateCompleted(taskID, userID, completed)
}

func (s *Service) DeleteTask(taskID int64) error {
	return s.repo.DeleteTask(taskID)
}

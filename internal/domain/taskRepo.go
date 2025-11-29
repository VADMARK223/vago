package domain

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetAllByUserID(ID uint) ([]Task, error)
	UpdateCompleted(taskID, userID uint, completed bool) error
}

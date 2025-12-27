package domain

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetAllByUserID(ID int64) ([]Task, error)
	UpdateCompleted(taskID, userID int64, completed bool) error
}

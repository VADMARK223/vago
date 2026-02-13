package domain

type TaskRepository interface {
	PostTask(name string, desc string, completed bool, userID int64) error
	GetAllByUserID(ID int64) ([]Task, error)
	UpdateCompleted(taskID, userID int64, completed bool) error
	DeleteTask(taskID int64) error
}

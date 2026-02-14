package domain

type UserRepository interface {
	CreateUser(user User) error
	DeleteUser(id int64) error
	GetByLogin(login string) (User, error)
	GetByID(id UserID) (User, error)
	GetByIDs(ids []int64) ([]User, error)
	GetAll() ([]User, error)
}

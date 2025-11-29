package domain

type UserRepository interface {
	CreateUser(user User) error
	DeleteUser(id uint) error
	GetByLogin(login string) (User, error)
	GetByID(id uint) (User, error)
	GetByIDs(ids []uint) ([]User, error)
	GetAll() ([]User, error)
}

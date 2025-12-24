package mocks

import (
	"vago/internal/domain"

	"github.com/stretchr/testify/mock"
)

// UserRepo - мок для domain.UserRepository
type UserRepo struct {
	mock.Mock
}

func (m *UserRepo) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepo) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepo) GetByLogin(login string) (domain.User, error) {
	args := m.Called(login)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepo) GetByID(id uint) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepo) GetByIDs(ids []uint) ([]domain.User, error) {
	args := m.Called(ids)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *UserRepo) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

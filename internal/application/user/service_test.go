package user

import (
	"errors"
	"testing"
	"vago/internal/application/user/mocks"
	"vago/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Test_Service_CreateUser_Success проверяет успешное создание пользователя
func Test_Service_CreateUser_Success(t *testing.T) {
	// Arrange (подготовка)
	repo := new(mocks.UserRepo)
	svc := NewService(repo, nil)

	dto := domain.DTO{
		Login:    "testuser",
		Password: "securePassword123",
		Email:    "test@example.com",
		Username: "Test User",
		Role:     domain.RoleUser,
		Color:    "#FF5733",
	}

	// Настраиваем мок: CreateUser должен быть вызван с любым User (пароль будет захеширован)
	repo.On("CreateUser", mock.MatchedBy(func(user domain.User) bool {
		// Проверяем, что пароль захеширован
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
		return err == nil &&
			user.Login == dto.Login &&
			user.Email == dto.Email &&
			user.Username == dto.Username &&
			user.Role == dto.Role &&
			user.Color == dto.Color
	})).Return(nil)

	// Act (действие)
	err := svc.CreateUser(dto)

	// Assert (проверка)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

// Test_Service_CreateUser_RepositoryError проверяет обработку ошибки репозитория
func Test_Service_CreateUser_RepositoryError(t *testing.T) {
	// Arrange
	repo := new(mocks.UserRepo)
	svc := NewService(repo, nil)

	dto := domain.DTO{
		Login:    "testuser",
		Password: "password123",
		Email:    "test@example.com",
		Username: "Test User",
		Role:     domain.RoleUser,
		Color:    "#FF5733",
	}

	expectedErr := errors.New("пользователь с таким логином уже существует")
	repo.On("CreateUser", mock.Anything).Return(expectedErr)

	// Act
	err := svc.CreateUser(dto)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	repo.AssertExpectations(t)
}

// Test_Service_CreateUser_PasswordHashing проверяет, что пароль действительно хешируется
func Test_Service_CreateUser_PasswordHashing(t *testing.T) {
	// Arrange
	repo := new(mocks.UserRepo)
	svc := NewService(repo, nil)

	plainPassword := "mySecretPassword123"
	dto := domain.DTO{
		Login:    "testuser",
		Password: plainPassword,
		Email:    "test@example.com",
		Username: "Test User",
		Role:     domain.RoleUser,
		Color:    "#FF5733",
	}

	var capturedUser domain.User
	repo.On("CreateUser", mock.Anything).Run(func(args mock.Arguments) {
		capturedUser = args.Get(0).(domain.User)
	}).Return(nil)

	// Act
	err := svc.CreateUser(dto)

	// Assert
	assert.NoError(t, err)

	// Проверяем, что пароль НЕ остался в открытом виде
	assert.NotEqual(t, plainPassword, capturedUser.Password, "Пароль должен быть захеширован")

	// Проверяем, что хеш можно верифицировать
	err = bcrypt.CompareHashAndPassword([]byte(capturedUser.Password), []byte(plainPassword))
	assert.NoError(t, err, "Хеш пароля должен совпадать с исходным паролем")

	repo.AssertExpectations(t)
}

// Test_Service_CreateUser_DifferentRoles тестирует создание пользователей с разными ролями
func Test_Service_CreateUser_DifferentRoles(t *testing.T) {
	tests := []struct {
		name string
		role domain.Role
	}{
		{"User role", domain.RoleUser},
		{"Moderator role", domain.RoleModerator},
		{"Admin role", domain.RoleAdmin},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			repo := new(mocks.UserRepo)
			svc := NewService(repo, nil)

			dto := domain.DTO{
				Login:    "testuser",
				Password: "password123",
				Email:    "test@example.com",
				Username: "Test User",
				Role:     tt.role,
				Color:    "#FF5733",
			}

			var capturedUser domain.User
			repo.On("CreateUser", mock.Anything).Run(func(args mock.Arguments) {
				capturedUser = args.Get(0).(domain.User)
			}).Return(nil)

			// Act
			err := svc.CreateUser(dto)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.role, capturedUser.Role)
			repo.AssertExpectations(t)
		})
	}
}

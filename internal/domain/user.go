package domain

import (
	"time"
)

type User struct {
	ID        int64
	Login     string
	Username  string
	Password  string
	Role      Role
	Color     string
	CreatedAt time.Time
}

func (u User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func New(login, password, username, color string, role Role) User {
	return User{
		Login:    login,
		Password: password,
		Username: username,
		Role:     role,
		Color:    color,
	}
}

type DTO struct {
	Login    string
	Password string
	Username string
	Role     Role
	Color    string
}

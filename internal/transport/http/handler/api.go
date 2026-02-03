package handler

import (
	"time"
	"vago/internal/domain"
)

type UsersApiDTO struct {
	Users []UserResponseDTO `json:"users"`
}

type UserResponseDTO struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
}

func userToResponse(u domain.User) UserResponseDTO {
	return UserResponseDTO{
		ID:        u.ID,
		Login:     u.Login,
		Username:  u.Username,
		Role:      string(u.Role),
		Color:     u.Color,
		CreatedAt: u.CreatedAt,
	}
}

func usersToResponse(users []domain.User) []UserResponseDTO {
	result := make([]UserResponseDTO, 0, len(users))
	for _, u := range users {
		result = append(result, userToResponse(u))
	}
	return result
}

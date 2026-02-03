package handler

import (
	"net/http"
	"time"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func OK[T any](c *gin.Context, message string, data T) {
	c.JSON(http.StatusOK, Response[T]{message, &data})
}

func OKNoData(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response[any]{
		Message: message,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response[any]{
		Message: message,
	})
}

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

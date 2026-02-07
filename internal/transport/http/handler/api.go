package handler

import (
	"time"
	"vago/internal/domain"
)

type UserApiDTO struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
}

type UsersApiDTO struct {
	Users []UserApiDTO `json:"users"`
}

func userToDTO(u domain.User) UserApiDTO {
	return UserApiDTO{
		ID:        u.ID,
		Login:     u.Login,
		Username:  u.Username,
		Role:      string(u.Role),
		Color:     u.Color,
		CreatedAt: u.CreatedAt,
	}
}

func usersToDTO(users []domain.User) []UserApiDTO {
	result := make([]UserApiDTO, 0, len(users))
	for _, u := range users {
		result = append(result, userToDTO(u))
	}
	return result
}

type TaskApiDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Completed   bool      `json:"completed"`
}

type TasksApiDTO struct {
	Tasks []TaskApiDTO `json:"tasks"`
}

func taskToDTO(t domain.Task) TaskApiDTO {
	return TaskApiDTO{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		Completed:   t.Completed,
	}
}

func tasksToDTO(tasks []domain.Task) []TaskApiDTO {
	result := make([]TaskApiDTO, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, taskToDTO(t))
	}

	return result
}

type PostTaskDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type UpdateTaskDTO struct {
	Completed bool `json:"completed"`
}

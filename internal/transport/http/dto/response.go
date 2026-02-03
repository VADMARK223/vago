package dto

import "vago/internal/domain"

type MeDTO struct {
	Username string      `json:"username"`
	Role     domain.Role `json:"role"`
}

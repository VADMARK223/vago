package error

import (
	"errors"
	"net/http"
	"vago/internal/domain"
)

func MapErrorToHTTP(err error) int {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

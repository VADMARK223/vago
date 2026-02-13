package handler

import (
	"errors"
	"net/http"
	"vago/internal/config/code"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

func TplWithMetaData(c *gin.Context, caption string) gin.H {
	td, exists := c.Get(code.TemplateData)
	if !exists {
		panic("TemplateData not found")
	}
	data := td.(gin.H)
	data[code.Caption] = caption
	path := c.Request.URL.Path
	data[code.Path] = path

	return data
}

func mapErrorToHTTP(err error) int {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

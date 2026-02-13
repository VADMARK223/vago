package response

import (
	"net/http"

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

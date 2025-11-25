package handler

import (
	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	ShowError(c, "Ошибка", "Страница не найдена.")
}

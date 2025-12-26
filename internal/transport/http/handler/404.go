package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	data := tplWithMetaData(c, "Страница не найдена")
	c.HTML(http.StatusInternalServerError, "404.html", data)
}

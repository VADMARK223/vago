package handler

import (
	"net/http"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	data := template.TplWithMetaData(c, "Страница не найдена")
	c.HTML(http.StatusInternalServerError, "404.html", data)
}

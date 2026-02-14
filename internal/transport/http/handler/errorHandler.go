package handler

import (
	"net/http"
	"vago/internal/config/code"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

func ShowError(c *gin.Context, capture string, err string) {
	data := template.TplWithMetaData(c, capture)
	data[code.Error] = err
	c.HTML(http.StatusInternalServerError, "error.html", data)
}

package handler

import (
	"net/http"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func ShowError(c *gin.Context, capture string, err string) {
	data := tplWithCapture(c, capture)
	data[code.Error] = err
	c.HTML(http.StatusInternalServerError, "error.html", data)
}

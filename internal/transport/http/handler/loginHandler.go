package handler

import (
	"net/http"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	data := tplWithMetaData(c, "Вход")
	if errVal, exists := c.Get(code.Error); exists {
		data[code.Error] = errVal
	}

	c.HTML(http.StatusOK, "sign_in.html", data)
}

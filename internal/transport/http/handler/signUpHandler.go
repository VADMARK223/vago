package handler

import (
	"net/http"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func ShowSignup(c *gin.Context) {
	data := tplWithMetaData(c, "Регистрация")
	if errVal, exists := c.Get(code.Error); exists {
		data[code.Error] = errVal
	}
	c.HTML(http.StatusOK, "sign_up.html", data)
}

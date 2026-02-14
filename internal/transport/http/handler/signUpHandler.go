package handler

import (
	"net/http"
	"vago/internal/config/code"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

func ShowSignup(c *gin.Context) {
	data := template.TplWithMetaData(c, "Регистрация")
	if errVal, exists := c.Get(code.Error); exists {
		data[code.Error] = errVal
	}
	c.HTML(http.StatusOK, "sign_up.html", data)
}

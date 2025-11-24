package handler

import (
	"net/http"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func Grpc(c *gin.Context) {
	token, _ := c.Cookie(code.VagoToken)
	data := tplWithCapture(c, "Test gRPC")
	data[code.VagoToken] = token
	c.HTML(http.StatusOK, "grpc-test.html", data)
}

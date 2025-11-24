package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Grpc(c *gin.Context) {
	data := tplWithCapture(c, "Проверка gRPC")
	c.HTML(http.StatusOK, "grpc-test.html", data)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "0.15.1"

func ShowIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		capture := "Vago портал"
		if gin.Mode() == gin.DebugMode {
			capture += " (debug)"
		}
		data := tplWithCapture(c, capture)
		data["version"] = version
		c.HTML(http.StatusOK, "index.html", data)
	}
}

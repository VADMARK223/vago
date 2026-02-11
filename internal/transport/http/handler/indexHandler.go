package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.16.0"

func ShowIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		caption := "Портал по изучению Golang"
		if gin.Mode() == gin.DebugMode {
			caption += " (debug)"
		}
		data := tplWithMetaData(c, caption)
		data["version"] = version
		c.HTML(http.StatusOK, "index.html", data)
	}
}

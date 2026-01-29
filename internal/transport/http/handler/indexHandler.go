package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

func ShowIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		capture := "Портал по изучению Golang"
		if gin.Mode() == gin.DebugMode {
			capture += " (debug)"
		}
		data := tplWithMetaData(c, capture)
		data["version"] = version
		c.HTML(http.StatusOK, "index.html", data)
	}
}

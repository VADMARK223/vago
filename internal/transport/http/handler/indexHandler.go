package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.18.1"

func ShowIndex(c *gin.Context) {
	caption := "Портал по изучению Golang"
	if gin.Mode() == gin.DebugMode {
		caption += " (debug)"
	}
	data := TplWithMetaData(c, caption)
	data["version"] = version
	c.HTML(http.StatusOK, "index.html", data)
}

package handler

import (
	"net/http"
	"vago/internal/transport/http/shared/template"

	"github.com/gin-gonic/gin"
)

const version = "1.26.1"

func ShowIndex(c *gin.Context) {
	caption := "Портал по изучению Golang"
	if gin.Mode() == gin.DebugMode {
		caption += " (debug)"
	}
	data := template.TplWithMetaData(c, caption)
	data["version"] = version
	c.HTML(http.StatusOK, "index.html", data)
}

package handler

import (
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func tplWithMetaData(c *gin.Context, capture string) gin.H {
	td, exists := c.Get(code.TemplateData)
	if !exists {
		panic("TemplateData not found")
	}
	data := td.(gin.H)
	data[code.Capture] = capture
	path := c.Request.URL.Path
	data[code.Path] = path

	return data
}

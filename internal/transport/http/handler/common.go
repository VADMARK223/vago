package handler

import (
	"fmt"
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
	fmt.Printf("➡️ \033[93m%s: \033[92m%v\033[0m\n", "path", path)
	data[code.Path] = path

	return data
}

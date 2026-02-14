package template

import (
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func TplWithMetaData(c *gin.Context, caption string) gin.H {
	td, exists := c.Get(code.TemplateData)

	if !exists {
		panic("TemplateData not found")
	}
	data := td.(gin.H)
	data[code.Caption] = caption
	path := c.Request.URL.Path
	data[code.Path] = path

	return data
}

func BaseAdminData(c *gin.Context, name string) gin.H {
	data := TplWithMetaData(c, "Админка ("+name+")")
	return data
}

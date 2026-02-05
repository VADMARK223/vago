package middleware

import (
	"vago/internal/config/code"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

func TemplateContext(c *gin.Context) {
	result := gin.H{
		code.Username: "Гость",
	}

	if contextUser, ok := c.Get(code.CurrentUser); ok {
		u := contextUser.(domain.User)

		result[code.UserId] = u.ID
		result[code.Login] = u.Login
		result[code.Username] = u.Username
		result[code.Role] = u.Role
		result[code.IsAdminModerator] = u.IsAdminOrModerator()
	}

	c.Set(code.TemplateData, result)

	c.Next()
}

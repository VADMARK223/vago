package middleware

import (
	"net/http"
	"vago/internal/config/code"
	"vago/internal/config/route"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuthAndRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get(code.UserId); !ok {
			session := sessions.Default(c)
			session.Set(code.RedirectTo, c.Request.URL.Path)
			_ = session.Save()
			c.Redirect(http.StatusFound, route.Login)
			c.Abort()
			return
		}

		c.Next()
	}
}

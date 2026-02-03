package middleware

import (
	"net/http"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/transport/http/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireAuthAndRedirect(c *gin.Context) {
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

func RequireAuthApi(c *gin.Context) {
	if _, ok := c.Get(code.UserId); !ok {
		api.Error(c, http.StatusUnauthorized, "Пользователь не аутентифицирован.")
		c.Abort()
		return
	}

	c.Next()
}

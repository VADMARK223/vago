package middleware

import (
	"net/http"
	"time"
	"vago/internal/app"
	"vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/domain"

	"github.com/gin-gonic/gin"
)

func LoadUserContext(svc *user.Service, cache *app.LocalCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidVal, exists := c.Get(code.UserId)
		if !exists {
			c.Next()
			return
		}

		if _, exists := c.Get(code.CurrentUser); exists {
			c.Next()
			return
		}

		userID := uidVal.(uint)

		if cached, ok := cache.Get(userID); ok {
			c.Set(code.CurrentUser, cached.(domain.User))
			c.Next()
			return
		}

		u, err := svc.GetByID(userID)
		if err != nil {
			domain.ClearTokenCookies(c)
			c.Redirect(http.StatusFound, route.Login)
			c.Abort()
		}

		cache.Set(userID, u, time.Minute*5)
		c.Set(code.CurrentUser, u)

		c.Next()
	}
}

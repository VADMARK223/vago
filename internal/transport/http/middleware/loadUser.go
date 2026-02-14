package middleware

import (
	"net/http"
	"time"
	"vago/internal/app"
	"vago/internal/application/user"
	"vago/internal/config/code"
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

		userID, ok := uidVal.(int64)
		if !ok {
			c.Next()
			return
		}

		if cached, ok := cache.Get(userID); ok {
			c.Set(code.CurrentUser, cached.(domain.User))
			c.Next()
			return
		}

		u, err := svc.GetByID(domain.UserID(userID))
		if err != nil {
			domain.ClearTokenCookies(c)
			//c.Redirect(http.StatusFound, route.Login)
			//c.Abort()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		cache.Set(userID, u, time.Minute*5)
		c.Set(code.CurrentUser, u)

		c.Next()
	}
}

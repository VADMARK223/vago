package middleware

import (
	"time"
	"vago/internal/config/code"
	"vago/internal/domain/auth"
	"vago/internal/infra/token"

	"github.com/gin-gonic/gin"
)

func CheckJWT(provider *token.JWTProvider, refreshTTL int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie(code.VagoToken)
		if err == nil && tokenStr != "" {
			claims, err := provider.ParseToken(tokenStr)
			if err == nil && (claims.ExpiresAt == nil || claims.ExpiresAt.Time.After(time.Now())) {
				setAuth(c, claims.UserID(), claims.Role, claims.ExpiresAt.Time)
				c.Next()
				return
			}
		}

		if tryRefresh(c, refreshTTL, provider) {
			c.Next()
			return
		}

		c.Next()
	}
}

func tryRefresh(c *gin.Context, refreshTTL int, provider *token.JWTProvider) bool {
	refreshStr, err := c.Cookie(code.VagoRefreshToken)
	if err != nil || refreshStr == "" {
		return false
	}

	refreshClaims, err := provider.ParseToken(refreshStr)
	if err != nil || (refreshClaims.ExpiresAt != nil && refreshClaims.ExpiresAt.Time.Before(time.Now())) {
		return false
	}

	newAccess, err := provider.CreateToken(refreshClaims.UserID(), refreshClaims.Role, true)
	if err != nil {
		return false
	}

	claims, _ := provider.ParseToken(newAccess)
	auth.SetCookie(c, code.VagoToken, newAccess, refreshTTL)

	setAuth(c, refreshClaims.UserID(), refreshClaims.Role, claims.ExpiresAt.Time)
	return true
}

func setAuth(c *gin.Context, id uint, role string, accessExpTime time.Time) {
	c.Set(code.UserId, id)
	c.Set(code.Role, role)
	c.Set(code.AccessExpTime, accessExpTime)
}

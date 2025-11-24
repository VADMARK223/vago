package auth

import (
	"net/http"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

func SetTokenCookies(c *gin.Context, tokens *TokenPair, refreshTTL int) {
	SetCookie(c, code.VagoToken, tokens.AccessToken, refreshTTL, false)
	SetCookie(c, code.VagoRefreshToken, tokens.RefreshToken, refreshTTL, true)
}

func SetCookie(c *gin.Context, name, value string, maxAge int, httpOnly bool) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     "/",
		Domain:   "",
		Secure:   false, // Cookie отправляется даже по HTTP (Надо поменять в production) Защита от MITM
		HttpOnly: httpOnly,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookieData(cookie)
}

func ClearTokenCookies(c *gin.Context) {
	clearCookie(c, code.VagoToken)
	clearCookie(c, code.VagoRefreshToken)
}

func clearCookie(c *gin.Context, name string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

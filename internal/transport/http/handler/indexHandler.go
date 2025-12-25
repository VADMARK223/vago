package handler

import (
	"fmt"
	"net/http"
	"time"
	"vago/internal/config/code"
	"vago/internal/infra/token"
	"vago/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

const version = "0.12.0"

func ShowIndex(provider *token.JWTProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		capture := "Vago портал"
		if gin.Mode() == gin.DebugMode {
			capture += " (debug)"
		}
		data := tplWithCapture(c, capture)

		updateTokenInfo(c, data)
		updateRefreshTokenInfo(c, data, provider)
		data["version"] = version
		c.HTML(http.StatusOK, "index.html", data)
	}
}

func updateTokenInfo(c *gin.Context, data gin.H) {
	data[code.TokenStatus] = "❌ информации о токене нет в контексте"
	data[code.TokenExpireAt] = "-"

	info, ok := middleware.TokenInfo(c)
	if !ok {
		return
	}

	data[code.TokenStatus] = "✅"
	data[code.TokenExpireAt] = fmt.Sprintf("%s (через %s)", info.Exp.Format("02.01.2006 15:04:05"), info.Remaining.String())
}

func updateRefreshTokenInfo(c *gin.Context, data gin.H, provider *token.JWTProvider) {
	data[code.RefreshTokenStatus] = "✅"
	data[code.RefreshTokenExpireAt] = "-"

	tokenStr, errTokenCookie := c.Cookie(code.VagoRefreshToken)

	if errTokenCookie != nil {
		data[code.RefreshTokenStatus] = "❌" + errTokenCookie.Error()
		return
	}

	claims, err := provider.ParseToken(tokenStr)
	if err != nil {
		data[code.RefreshTokenStatus] = "❌" + err.Error()
		return
	}

	expTime := claims.ExpiresAt.Time
	remaining := time.Until(expTime).Truncate(time.Second)
	data[code.RefreshTokenExpireAt] = fmt.Sprintf("%s (через %s)", expTime.Format("02.01.2006 15:04:05"), remaining.String())
}

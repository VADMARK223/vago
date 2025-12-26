package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"vago/internal/application/chat"
	"vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/infra/token"
	"vago/internal/transport/http/middleware"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	provider *token.JWTProvider
	userSvc  *user.Service
	chatSvc  *chat.Service
}

func NewAdminHandler(provider *token.JWTProvider, userSvc *user.Service, chatSvc *chat.Service) *AdminHandler {
	return &AdminHandler{
		provider: provider,
		userSvc:  userSvc,
		chatSvc:  chatSvc,
	}
}

func (h *AdminHandler) ShowAdmin(c *gin.Context) {
	c.Redirect(http.StatusFound, route.Admin+route.Messages)
}

func (h *AdminHandler) ShowUser(c *gin.Context) {
	data := baseAdminData(c, "Пользователь")
	updateTokenInfo(c, data)
	updateRefreshTokenInfo(c, data, h.provider)
	data["Active"] = "user"
	c.HTML(http.StatusOK, "admin/layout", data)
}

func (h *AdminHandler) ShowUsers(c *gin.Context) {
	users, err := h.userSvc.GetAll()
	if err != nil {
		ShowError(c, "Ошибка загрузки пользователя", err.Error())
		return
	}

	data := baseAdminData(c, "Пользователи")
	data["Users"] = users
	data["Active"] = "users"
	c.HTML(http.StatusOK, "admin/layout", data)
}

func (h *AdminHandler) ShowMessages(c *gin.Context) {
	all, err := h.chatSvc.MessagesDTO(context.Background())
	if err != nil {
		ShowError(c, "Ошибка получения списка сообщений", err.Error())
		return
	}

	data := baseAdminData(c, "Сообщения")
	data[code.Messages] = all
	data[code.MessagesCount] = len(all)
	data["Active"] = "messages"
	c.HTML(http.StatusOK, "admin/layout", data)
}

func (h *AdminHandler) ShowGrpc(c *gin.Context) {
	data := baseAdminData(c, "Тест gRPC")
	data["Active"] = "grpc"
	c.HTML(http.StatusOK, "admin/layout", data)
}

func baseAdminData(c *gin.Context, name string) gin.H {
	data := tplWithCapture(c, "Админка ("+name+")")
	return data
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

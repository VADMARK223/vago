package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"vago/internal/app"
	"vago/internal/chat/chatApp"
	"vago/internal/chat/domain"
	"vago/internal/config/code"
	"vago/internal/infra/token"

	"vago/internal/trasport/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func ShowChat(port string, service *chatApp.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		all, err := service.LastMessages(context.Background())
		if err != nil {
			ShowError(c, "Ошибка получения списка сообщений", err.Error())
			return
		}

		data := tplWithCapture(c, "Чат")
		data[code.Port] = port
		dtos := make([]domain.MessageDTO, 0, len(all))
		for _, m := range all {
			dtos = append(dtos, domain.MessageDTO{
				//ID:     m.ID,
				Author: m.Author(),
				Body:   m.Body(),
				SentAt: m.SentAt(),
				Type:   "message", // TODO: пофиксякать
			})
		}
		jsonBytes, _ := json.Marshal(dtos)
		data["messages_json"] = string(jsonBytes)
		app.Dump("Messages", data["messages_json"])
		c.HTML(http.StatusOK, "chat.html", data)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Разрешаем любые origins (можно ужесточить)
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeSW(hub *ws.Hub, log *zap.SugaredLogger, provider *token.JWTProvider, svc *chatApp.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		// Если токена нет в query параметре, пробуем взять из кук
		if tokenStr == "" {
			var err error
			tokenStr, err = c.Cookie(code.VagoToken)
			if err != nil || tokenStr == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
				return
			}
		}
		log.Infow("ServeSW", "tokenStr", tokenStr)

		claims, err := provider.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Errorw("Upgrader error", "error", err)
			return
		}

		client := ws.NewClient(conn, hub, claims.UserID(), log, svc)
		hub.Register <- client

		go client.OutgoingLoop()
		go client.IncomingLoop()
	}
}

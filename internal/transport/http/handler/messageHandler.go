package handler

import (
	"context"
	"net/http"
	"strconv"
	"vago/internal/application/chat"
	"vago/internal/application/user"
	"vago/internal/config/code"
	"vago/internal/config/route"
	"vago/internal/domain"
	"vago/internal/transport/http/api"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	chatSvc *chat.Service
	userSvc *user.Service
}

func NewMessageHandler(chatSvc *chat.Service, userSvc *user.Service) *MessageHandler {
	return &MessageHandler{chatSvc: chatSvc, userSvc: userSvc}
}

func (h *MessageHandler) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, errRole := c.Get(code.Role)
		if !errRole {
			api.Error(c, http.StatusBadRequest, "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			api.Error(c, http.StatusForbidden, "У вас нет прав на удаление пользователей")
			return
		}

		parseUint, parseUintErr := strconv.ParseInt(c.Param("id"), 10, 64)
		if parseUintErr != nil {
			ShowError(c, "Ошибка конвертации идентификатора", parseUintErr.Error())
			return
		}

		errDelete := h.chatSvc.DeleteMessage(parseUint)

		if errDelete != nil {
			ShowError(c, "Ошибка удаления сообщения", errDelete.Error())
			return
		}

		c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
	}
}

func (h *MessageHandler) DeleteAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, errRole := c.Get(code.Role)
		if !errRole {
			ShowError(c, "Ошибка удаления всех сообщений", "Роль пользователя неизвестна")
			return
		}

		if role != domain.RoleAdmin {
			ShowError(c, "Ошибка удаления всех сообщений", "У вас нет прав на удаление пользователей")
			return
		}

		err := h.chatSvc.DeleteAll()
		if err != nil {
			ShowError(c, "Ошибка удаления всех сообщений", err.Error())
		}
		c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
	}
}

func (h *MessageHandler) AddMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		createDTO := chat.MessageCreateDTO{AuthorID: 1, Body: "Test", MessageType: "text"}
		_, err := h.chatSvc.CreateMessage(context.Background(), createDTO)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error add message")
			return
		}

		c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
	}
}

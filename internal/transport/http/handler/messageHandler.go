package handler

import (
	"context"
	"net/http"
	"strconv"
	"vago/internal/application/chat"
	"vago/internal/application/user"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	chatSvc *chat.Service
	userSvc *user.Service
}

func NewMessageHandler(chatSvc *chat.Service, userSvc *user.Service) *MessageHandler {
	return &MessageHandler{chatSvc: chatSvc, userSvc: userSvc}
}

func (h *MessageHandler) ShowMessages() func(c *gin.Context) {
	return func(c *gin.Context) {
		all, err := h.chatSvc.MessagesDTO(context.Background())
		if err != nil {
			ShowError(c, "Ошибка получения списка сообщений", err.Error())
			return
		}

		data := tplWithCapture(c, "Сообщения")
		data[code.Messages] = all
		data[code.MessagesCount] = len(all)

		c.HTML(http.StatusOK, "messages.html", data)
	}
}

func (h *MessageHandler) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		parseUint, parseUintErr := strconv.ParseUint(c.Param("id"), 10, 32)
		if parseUintErr != nil {
			ShowError(c, "Ошибка конвертации идентификатора", parseUintErr.Error())
			return
		}

		errDelete := h.chatSvc.DeleteMessage(uint(parseUint))

		if errDelete != nil {
			ShowError(c, "Ошибка удаления сообщения", errDelete.Error())
			return
		}

		c.Redirect(http.StatusSeeOther, "/messages")
	}
}

func (h *MessageHandler) DeleteAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := h.chatSvc.DeleteAll()
		if err != nil {
			ShowError(c, "Ошибка удаления всех", err.Error())
		}
		c.Redirect(http.StatusSeeOther, "/messages")
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

		c.Redirect(http.StatusSeeOther, "/messages")
	}
}

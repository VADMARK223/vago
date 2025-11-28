package handler

import (
	"context"
	"net/http"
	"strconv"
	"vago/internal/app"
	"vago/internal/chat/chatApp"
	"vago/internal/chat/domain"
	"vago/internal/config/code"
	"vago/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	chatSvc *chatApp.Service
	userSvc *user.Service
}

func NewMessageHandler(chatSvc *chatApp.Service, userSvc *user.Service) *MessageHandler {
	return &MessageHandler{chatSvc: chatSvc, userSvc: userSvc}
}

func (h *MessageHandler) ShowMessages() func(c *gin.Context) {
	return func(c *gin.Context) {
		all, err := h.chatSvc.Messages(context.Background())
		if err != nil {
			ShowError(c, "Ошибка получения списка сообщений", err.Error())
			return
		}

		data := tplWithCapture(c, "Сообщения")
		messagesCount := len(all)
		dtos := make([]domain.MessageDTO, 0, messagesCount)
		for _, m := range all {
			u, errUser := h.userSvc.GetByID(uint(m.Author()))
			var username string
			if errUser != nil {
				username = "Неизвестно"
			} else {
				username = u.Username
			}
			dtos = append(dtos, m.ToDTO(username))
		}

		data[code.Messages] = dtos
		data["messages_count"] = messagesCount

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
		err := h.chatSvc.SendMessage(context.Background(), 1, "Hello")
		if err != nil {
			app.Dump("Error add message", err)
			c.String(http.StatusInternalServerError, "Error add message")
			return
		}

		//c.Redirect(http.StatusFound, c.Request.Referer())
		c.Redirect(http.StatusSeeOther, "/messages")
	}
}

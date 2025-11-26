package handler

import (
	"context"
	"net/http"
	"strconv"
	"vago/internal/app"
	"vago/internal/chat/chatApp"
	"vago/internal/config/code"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service *chatApp.Service
}

func NewMessageHandler(service *chatApp.Service) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) ShowMessages() func(c *gin.Context) {
	return func(c *gin.Context) {
		renderMessagePage(c, h.service, "")
	}
}

func (h *MessageHandler) Delete() func(c *gin.Context) {
	return func(c *gin.Context) {
		parseUint, parseUintErr := strconv.ParseUint(c.Param("id"), 10, 32)
		if parseUintErr != nil {
			ShowError(c, "Ошибка конвертации идентификатора", parseUintErr.Error())
			return
		}

		app.Dump("ID", parseUint)
		app.Dump("parseUintErr", parseUintErr)
		errDelete := h.service.DeleteMessage(uint(parseUint))

		if errDelete != nil {
			ShowError(c, "Ошибка удаления сообщения", errDelete.Error())
			return
		}

		renderMessagePage(c, h.service, "")
	}
}

func (h *MessageHandler) DeleteAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := h.service.DeleteAll()
		if err != nil {
			ShowError(c, "Ошибка удаления", err.Error())
		}
		renderMessagePage(c, h.service, "")
	}
}

func renderMessagePage(c *gin.Context, service *chatApp.Service, _error string) {
	all, err := service.LastMessages(context.Background())
	if err != nil {
		ShowError(c, "Ошибка получения списка сообщений", err.Error())
		return
	}

	data := tplWithCapture(c, "Сообщения")
	data[code.Messages] = all

	c.HTML(http.StatusOK, "messages.html", data)
}

func (h *MessageHandler) AddMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := h.service.SendMessage(context.Background(), 1, "Hello")
		if err != nil {
			app.Dump("Error add message", err)
			c.String(http.StatusInternalServerError, "Error add message")
			return
		}

		renderMessagePage(c, h.service, "")
		//c.Redirect(http.StatusFound, c.Request.Referer())
		// c.Redirect(http.StatusSeeOther, "/chat")
	}
}

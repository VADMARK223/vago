package handler

import (
	"context"
	"net/http"
	"vago/internal/application/chat"
	"vago/internal/config/route"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	chatSvc *chat.Service
}

func NewMessageHandler(chatSvc *chat.Service) *MessageHandler {
	return &MessageHandler{chatSvc: chatSvc}
}

func (h *MessageHandler) AddMessage(c *gin.Context) {
	createDTO := chat.MessageCreateDTO{AuthorID: 1, Body: "Test", MessageType: "text"}
	_, err := h.chatSvc.SaveMessage(context.Background(), createDTO)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error add message")
		return
	}

	c.Redirect(http.StatusSeeOther, route.Admin+route.Messages)
}

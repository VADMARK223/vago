package message

import "github.com/gin-gonic/gin"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) GetAllMessages(c *gin.Context) {

}

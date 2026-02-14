package message

import (
	"vago/internal/application/message"

	"github.com/gin-gonic/gin"
)

type Loader struct {
	MessageSvc *message.Service
}

func (l Loader) Load(c *gin.Context) ([]message.WithUsername, error) {
	result, err := l.MessageSvc.ListMessagesWithAuthors(c.Request.Context())
	if err != nil {

		return []message.WithUsername{}, err

	}
	return result, nil
}

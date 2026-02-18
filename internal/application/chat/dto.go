package chat

import (
	"vago/internal/domain"
	//"vago/internal/transport/ws"
)

type MessageCreateDTO struct {
	AuthorID    domain.UserID `json:"author_id"`
	Body        string        `json:"body"`
	MessageType string        `json:"type"`
}

/*func New(id domain.UserID, packet ws.ClientPacket) MessageCreateDTO {
	return MessageCreateDTO{
		AuthorID:    id,
		Body:        packet.Text,
		MessageType: "text",
	}
}*/

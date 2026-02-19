package ws

import "encoding/json"

// Inbound формат сообщений, которые приходит от клиента
type Inbound struct {
	Type    string          `json:"type"`    // Типа сообщения
	Payload json.RawMessage `json:"payload"` // Содержание сообщения
}

type Outbound struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

const (
	TypeMessageSend = "message.send"
	TypeMessageNew  = "message.new"
	TypeError       = "error"
)

type MessageSendPayload struct {
	Text string `json:"text"`
}

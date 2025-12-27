package domain

import (
	"time"
)

type UserID int64
type Body string

type Message struct {
	ID          int64
	author      UserID
	body        Body
	sentAt      time.Time
	MessageType string
}

func NewMessage(author UserID, body Body, messageType string) *Message {
	return &Message{author: author, body: body, sentAt: time.Now(), MessageType: messageType}
}

func (m *Message) Author() UserID        { return m.author }
func (m *Message) Body() Body            { return m.body }
func (m *Message) SentAt() time.Time     { return m.sentAt }
func (m *Message) SetSentAt(t time.Time) { m.sentAt = t }

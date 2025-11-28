package domain

import (
	"errors"
	"time"
)

type UserID uint
type Body string

type Message struct {
	ID     uint
	author UserID
	body   Body
	sentAt time.Time
}

func New(id uint, author UserID, body Body, sentAt time.Time) (*Message, error) {
	if body == "" {
		return nil, errors.New("body is empty")
	}

	return &Message{ID: id, author: author, body: body, sentAt: sentAt}, nil
}

func (m *Message) Author() UserID    { return m.author }
func (m *Message) Body() Body        { return m.body }
func (m *Message) SentAt() time.Time { return m.sentAt }

type MessageDTO struct {
	Author UserID    `json:"author"`
	Body   Body      `json:"body"`
	SentAt time.Time `json:"sent_at"`

	Type string `json:"type"`
}

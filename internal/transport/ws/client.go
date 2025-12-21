package ws

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"vago/internal/app"
	"vago/internal/application/chat"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Conn       *websocket.Conn
	Hub        *Hub
	Send       chan []byte
	log        *zap.SugaredLogger
	UserID     uint
	messageSvc *chat.Service
}

type ClientPacket struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewClient(conn *websocket.Conn, hub *Hub, userID uint, log *zap.SugaredLogger, svc *chat.Service) *Client {
	client := &Client{
		Conn:       conn,
		Hub:        hub,
		Send:       make(chan []byte, 256),
		UserID:     userID,
		log:        log,
		messageSvc: svc,
	}
	return client
}

// IncomingLoop читает от клиента
func (c *Client) IncomingLoop() {
	defer func() {
		c.Hub.Unregister <- c
		_ = c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512)
	_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("WS read error:", err)
			break
		}

		var packet ClientPacket
		if errUnmarshal := json.Unmarshal(message, &packet); errUnmarshal != nil {
			log.Println("WS json error:", errUnmarshal)
			continue
		}

		c.log.Infow("Received message", "packet", packet)
		createDTO := chat.MessageCreateDTO{AuthorID: c.UserID, Body: packet.Text, MessageType: "text"}
		dto, errSendMessage := c.messageSvc.CreateMessage(context.Background(), createDTO)
		if errSendMessage != nil {
			// TODO: доделать.
			app.Dump("Send message error", errSendMessage)
			return
		}

		switch packet.Type {
		case "message":
			serverMsgBytes, _ := json.Marshal(dto)
			c.Hub.Broadcast <- serverMsgBytes
		default:
			c.log.Warnw("Unknown message type", "type", packet.Type)
		}
	}
}

// OutgoingLoop вытаскивает из канала сообщения, которые присылает менеджер сообщений.
func (c *Client) OutgoingLoop() {
	ticker := time.NewTicker(30 * time.Second)

	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Println("OutgoingLoop", "Channel closed:", err)
				}
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = w.Write(message)
			if err != nil {
				return
			}
			_ = w.Close()

		case <-ticker.C:
			// отправляем Ping
			_ = c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

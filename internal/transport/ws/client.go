package ws

import (
	"context"
	"encoding/json"
	"strings"
	"time"
	"vago/internal/application/chat"
	"vago/internal/domain"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Conn       *websocket.Conn    // Соединение
	Hub        *Hub               // Ссылка на менеджер
	Send       chan []byte        // Очередь исходящих сообщений (от сервера к клиенту)
	UserID     int64              // Идентификатор клиента
	messageSvc *chat.Service      // Сервис для работы с сообщениями (сохранение в БД)
	log        *zap.SugaredLogger // Логгер
}

func NewClient(conn *websocket.Conn, hub *Hub, userID int64, svc *chat.Service, log *zap.SugaredLogger) *Client {
	client := &Client{
		Conn:       conn,
		Hub:        hub,
		Send:       make(chan []byte, 256),
		UserID:     userID,
		messageSvc: svc,
		log:        log,
	}
	return client
}

// IncomingLoop читает от клиента
func (c *Client) IncomingLoop() {
	defer func() {
		select {
		case c.Hub.Unregister <- c: // При выходе из метода, отправит, что клиент покинул чат
		default: // Hub уже не работает или канал занят (тут не блокируемся)
		}

		_ = c.Conn.Close() // Закрываем сокет
	}()

	c.Conn.SetReadLimit(8 * 1024)                                // Максимум входящего сообщения 8 КБайт
	_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // Если 60 сек нет данных/понга, то read отвалится
	c.Conn.SetPongHandler(func(string) error {                   // На pong продлеваем read deadline еще на 60 сек
		_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for { // Бесконечны цикл чтения
		_, message, err := c.Conn.ReadMessage() // Читаем фрейм

		if err != nil { // Если есть ошибка чтения фрейма, то прерываем цикл чтения
			c.log.Infow("WS read error", "err", err)
			break
		}

		var in Inbound
		if err := json.Unmarshal(message, &in); err != nil {
			c.log.Errorw("Error unmarshal inbound", "err", err)
			continue
		}

		switch in.Type {
		case TypeMessageSend:
			var p MessageSendPayload
			if errUnmarshalPayload := json.Unmarshal(in.Payload, &p); errUnmarshalPayload != nil {
				c.log.Errorw("Error unmarshal payload", "err", errUnmarshalPayload)
				continue
			}

			text := strings.TrimSpace(p.Text)
			if text == "" {
				continue
			}

			c.log.Infow("Received message", "in", in)
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Тайм-аут на сохранение в БД
			// Сохраняем, сообщение в БД через сервис
			dto, err := c.messageSvc.SaveMessage(ctx, chat.MessageCreateDTO{
				AuthorID:    domain.UserID(c.UserID),
				Body:        text,
				MessageType: "text",
			})
			cancel()
			if err != nil {
				c.log.Errorw("Send message error", err)
				return
			}

			out := Outbound{Type: TypeMessageNew, Payload: dto}
			serverMsgBytes, err := json.Marshal(out)
			if err != nil {
				c.log.Errorw("Marshal error", "err", err)
				continue
			}

			select {
			case c.Hub.Broadcast <- serverMsgBytes: // Отправляем в канал общей рассылки другим клиентам
			default:
				// Hub не читает/остановлен/забит — не блокируемся
				c.log.Warn("Broadcast skipped: hub is not accepting messages")
			}

		default:
			c.log.Warnw("Unknown message type", "type", in.Type)
		}
	}
}

// OutgoingLoop вытаскивает из канала сообщения, которые присылает менеджер сообщений.
func (c *Client) OutgoingLoop() {
	ticker := time.NewTicker(30 * time.Second) // Каждые 30 секунд шлём ping

	defer func() {
		ticker.Stop()      // Остановка тикета
		_ = c.Conn.Close() // // Закрываем сокет
	}()

	for {
		select {
		case message, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)) // Устанавливаем deadline 10 сек
			if !ok {                                                      // Если канал закрыт
				err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}) // Отправляем close frame
				if err != nil {
					c.log.Infow("OutgoingLoop", "Channel closed", err)
				}
				return
			}

			writer, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = writer.Write(message) // Пишем байты
			if err != nil {
				return
			}
			_ = writer.Close()

		case <-ticker.C:
			// Отправляем Ping
			_ = c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

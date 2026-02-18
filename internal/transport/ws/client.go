package ws

import (
	"context"
	"encoding/json"
	"time"
	"vago/internal/app"
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

// ClientPacket формат сообщений, которые приходит от клиента
type ClientPacket struct {
	Type string `json:"type"` // Типа сообщения
	Text string `json:"text"` // Содержание сообщения
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

	// TODO: Увеличить, очень мало для реального сообщения
	c.Conn.SetReadLimit(512)                                     // Максимум входящего сообщения 512 байт
	_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // Если 60 сек нет данных/понга, то read отвалится
	c.Conn.SetPongHandler(func(string) error {                   // На pong продлеваем read deadline еще на 60 сек
		_ = c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for { // Бесконечны цикл чтения
		_, message, err := c.Conn.ReadMessage() // Читаем фрейм

		if err != nil { // Если есть ошибка чтения фрейма, то прерываем цикл чтения
			c.log.Infow("WS read error:", err)
			break
		}

		var packet ClientPacket
		if errUnmarshal := json.Unmarshal(message, &packet); errUnmarshal != nil {
			c.log.Errorw("WS json error:", errUnmarshal)
			continue
		}

		c.log.Infow("Received message", "packet", packet)
		dto, errSaveMessage := c.messageSvc.SaveMessage( // Сохраняем сообщение в БД через сервис
			context.Background(),
			chat.MessageCreateDTO{
				AuthorID:    domain.UserID(c.UserID),
				Body:        packet.Text,
				MessageType: "text",
			})

		if errSaveMessage != nil {
			// TODO: доделать.
			app.Dump("Send message error", errSaveMessage)
			return
		}

		switch packet.Type {
		case "message":
			serverMsgBytes, _ := json.Marshal(dto)

			select {
			case c.Hub.Broadcast <- serverMsgBytes: // Отправляем в канал общей рассылки другим клиентам
			default:
				// Hub не читает/остановлен/забит — не блокируемся
				c.log.Warn("Broadcast skipped: hub is not accepting messages")
			}

		default:
			c.log.Warnw("Unknown message type", "type", packet.Type)
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
					c.log.Infow("OutgoingLoop", "Channel closed:", err)
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

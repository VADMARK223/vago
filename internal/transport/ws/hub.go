package ws

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type Hub struct {
	Register   chan *Client       // Канал "подключился новый клиент"
	Unregister chan *Client       // Канал "клиент отключился/нужно убрать"
	Broadcast  chan []byte        // Канал "надо разослать всем одно сообщение"
	Clients    map[*Client]bool   // Текущее множество подключенных клиентов.
	log        *zap.SugaredLogger // Логгер
}

func NewHub(log *zap.SugaredLogger) *Hub {
	return &Hub{
		Register:   make(chan *Client, 64),
		Unregister: make(chan *Client, 64),
		Broadcast:  make(chan []byte, 256),
		Clients:    make(map[*Client]bool),
		log:        log,
	}
}

// Run единственная горутина, которая владеет Clients и изменяет его
func (h *Hub) Run(ctx context.Context) {
	h.log.Info("WebSocket Hub started")
	defer h.log.Info("WebSocket Hub stopped")
Loop:
	for {
		select {
		case <-ctx.Done(): // Выключения сервера
			h.log.Infow("Hub shutting down, disconnecting clients", "count", len(h.Clients))
			for client := range h.Clients {
				close(client.Send)        // Закрываем канал отправки клиенту
				delete(h.Clients, client) // Удаляем из карты
			}
			return // Прерываем бесконечный цикл

		case c := <-h.Register: // Пришел новый клиент
			h.log.Info("Registered client")
			// Собираем snapshot пользователей (до добавления нового)
			users := make([]map[string]any, 0, len(h.Clients))
			for cl := range h.Clients {
				users = append(users, map[string]any{
					"userId":   cl.UserID,
					"username": cl.Username,
				})
			}
			snapshot := Outbound{
				Type: "users.snapshot",
				Payload: map[string]any{
					"users": users,
				},
			}
			bSnapshot, err := json.Marshal(snapshot)
			if err != nil {
				h.log.Errorw("Marshal snapshot error", "err", err)
				close(c.Send)
				continue Loop
			}

			// Пытаемся отправить snapshot новому клиенту
			select {
			case c.Send <- bSnapshot:
			default:
				// Клиент ещё не готов/буфер забит — просто не регистрируем
				close(c.Send)
				continue Loop
			}

			h.Clients[c] = true // Добавляем клиента в карту Clients

			// Рассылаем всем что пользователь присоединился к чату
			joined := Outbound{
				Type: "user.joined",
				Payload: map[string]any{
					"userId":   c.UserID,
					"username": c.Username,
				},
			}
			b, err := json.Marshal(joined)
			if err != nil {
				h.log.Errorw("Marshal joined error", "err", err)
				continue Loop
			}

			for cl := range h.Clients {
				if cl == c {
					continue // Исключаем только что подключившегося
				}
				select {
				case cl.Send <- b:
				default:
					delete(h.Clients, cl)
					close(cl.Send)
				}
			}

		case c := <-h.Unregister: // Клиент ушел
			h.log.Info("Unregistered client")
			if _, ok := h.Clients[c]; ok { // Проходимся по карте
				delete(h.Clients, c) // Удаляем клиента из карты
				close(c.Send)        // закрываем канал (Сигнал для OutgoingLoop, что порка завершиться)

				left := Outbound{
					Type:    "user.left",
					Payload: map[string]any{"userId": c.UserID},
				}
				b, _ := json.Marshal(left)
				// Разослать всем оставшимся пользователям, что участник вышел
				for cl := range h.Clients {
					select {
					case cl.Send <- b:
					default:
						delete(h.Clients, cl)
						close(cl.Send)
					}
				}
			}

		case msg := <-h.Broadcast: // Общая рассылка
			for c := range h.Clients { // Проходим по карте
				select {
				case c.Send <- msg: // Клиентам пытаемся отослать сообщение
				default: // Если клиент не успевает читать (буфер Send забит)
					delete(h.Clients, c) // Клиент не успевает читать, отключаем его (Жесткий кик медленного клиента)
					close(c.Send)
				}
			}
		}
	}
}

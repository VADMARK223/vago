package ws

import (
	"context"

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
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		Clients:    make(map[*Client]bool),
		log:        log,
	}
}

// Run единственная горутина, которая владеет Clients и изменяет его
func (h *Hub) Run(ctx context.Context) {
	h.log.Info("WebSocket Hub started")
	defer h.log.Info("WebSocket Hub stopped")

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
			h.Clients[c] = true // Добавляем клиента в карту Clients

		case c := <-h.Unregister: // Клиент ушел
			h.log.Info("Unregistered client")
			if _, ok := h.Clients[c]; ok { // Проходимся по карте
				delete(h.Clients, c) // Удаляем клиента из карты
				close(c.Send)        // закрываем канал (Сигнал для OutgoingLoop, что порка завершиться)
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

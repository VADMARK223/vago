package ws

import (
	"context"

	"go.uber.org/zap"
)

type Hub struct {
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	Clients    map[*Client]bool
	log        *zap.SugaredLogger
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

// Run запускает Hub и слушает ctx.Done() для graceful shutdown
func (h *Hub) Run(ctx context.Context) {
	h.log.Info("WebSocket Hub started")
	defer h.log.Info("WebSocket Hub stopped")

	for {
		select {
		case <-ctx.Done():
			h.log.Infow("Hub shutting down, disconnecting clients", "count", len(h.Clients))
			for client := range h.Clients {
				close(client.Send)
				delete(h.Clients, client)
			}
			return

		case c := <-h.Register:
			h.log.Info("Registered client")
			h.Clients[c] = true

		case c := <-h.Unregister:
			h.log.Info("Unregistered client")
			if _, ok := h.Clients[c]; ok {
				delete(h.Clients, c)
				close(c.Send)
			}

		case msg := <-h.Broadcast:
			for c := range h.Clients {
				select {
				case c.Send <- msg:
				default:
					// Клиент не успевает читать, отключаем его
					delete(h.Clients, c)
					close(c.Send)
				}
			}
		}
	}
}

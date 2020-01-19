package tcp

import (
	"fmt"
)

type Hub struct {
	clients map[*HubClient]bool
	broadcast chan []byte
	register chan *HubClient
	unregister chan *HubClient
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *HubClient),
		unregister: make(chan *HubClient),
		clients:    make(map[*HubClient]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("hub register")
			h.clients[client] = true
			go client.writePump()
			go client.readPump()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Println("hub broadcast")
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
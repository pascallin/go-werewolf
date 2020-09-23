package tcp

import (
	"fmt"
	"io"
)

type Hub struct {
	clients    	map[*HubClient]bool
	broadcast  	chan []byte
	register   	chan *HubClient
	unregister 	chan *HubClient
	writer 		io.Writer
}

func newHub(writer io.Writer) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *HubClient),
		unregister: make(chan *HubClient),
		clients:    make(map[*HubClient]bool),
		writer: 	writer,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Fprintf(h.writer,"hub register: " + client.conn.RemoteAddr().String() + "\n")
			h.clients[client] = true
			go client.writePump()
			go client.readPump()
		case client := <-h.unregister:
			fmt.Fprintf(h.writer,"hub unregister: " + client.conn.RemoteAddr().String() + "\n")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Fprintf(h.writer,"hub broadcast: " + string(message) + "\n")
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

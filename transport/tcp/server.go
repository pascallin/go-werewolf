package tcp

import (
	"log"
	"net"
)
type Server struct {
	hub *Hub
	listener net.Listener
}

func NewServer() *Server {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	server := &Server{
		newHub(),
		listener,
	}
	go server.hub.run()
	go server.listen()

	return server
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
		}

		// create server side client
		client := &HubClient{
			hub: s.hub,
			conn: conn,
			send: make(chan []byte, 256),
		}
		// user register
		s.hub.register <- client
	}
}
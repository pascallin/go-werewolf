package tcp

import (
	"io"
	"log"
	"net"
)

type Server struct {
	hub      	*Hub
}

func NewServer(writer io.Writer) *Server {
	server := &Server{
		newHub(writer),
	}

	return server
}

func (s *Server) Listen(port string) {
	listener, err := net.Listen("tcp", "localhost:" + port)
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}
	go s.hub.run()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
			continue
		}

		// create server side client
		client := &HubClient{
			hub:  s.hub,
			conn: conn,
			send: make(chan []byte, 256),
		}
		// user register
		s.hub.register <- client
	}
}

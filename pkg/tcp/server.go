package tcp

import (
	"bufio"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"net"
)

type Server struct {
	clients    	map[*ConnClient]bool
	Register   	chan *ConnClient
	Unregister 	chan *ConnClient
	Receiver	chan string
}

func NewServer() *Server {
	server := &Server{
		map[*ConnClient]bool{},
		make(chan *ConnClient),
		make(chan *ConnClient),
		make(chan string),
	}

	return server
}

func (s *Server) Listen(port string) {
	listener, err := net.Listen("tcp", "localhost:" + port)
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}
	go s.listenRegister(listener)
}

func (s *Server) listenRegister(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error", err)
			continue
		}
		s.registerConnClient(conn)
	}
}

func (s *Server) registerConnClient(conn net.Conn) {
	// create server side client
	client := &ConnClient{
		ID:		uuid.NewV4(),
		conn: 	conn,
	}
	s.clients[client] = true

	// user register notice
	s.Register <- client

	defer func() {
		client.close()
		s.unregisterConnClient(client)
	}()
	for {
		message, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil || err == io.EOF {
			continue
		}
		s.Receiver <- message
	}
}

func (s *Server) unregisterConnClient(client *ConnClient) {
	if _, ok := s.clients[client]; ok {
		delete(s.clients, client)
	}
	// user unregister notice
	s.Unregister <- client
}

func (s *Server) BroadcastMessage(message string) {
	for client := range s.clients {
		client.send(message)
	}
}

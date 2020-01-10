package tcpsocket

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var hub = newHub()

func handleConnection(conn net.Conn) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		log.Println("client left..")
		conn.Close()

		return
	}

	message := string(bufferBytes)
	clientAddr := conn.RemoteAddr().String()
	response := fmt.Sprintf(message + " from " + clientAddr + "\n")

	fmt.Println("tcp server handleConnection")
	hub.broadcast <- []byte(response)

	go handleConnection(conn)
}

func NewServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	go hub.run()

	for {
		conn, err := listener.Accept()

		// user register
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
		hub.register <- client

		if err != nil {
			log.Fatal("tcp server accept error", err)
		}

		go handleConnection(conn)
	}
}
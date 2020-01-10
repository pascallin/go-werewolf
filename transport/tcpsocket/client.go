package tcpsocket

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	hub *Hub
	conn net.Conn
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// TODO
				return
			}
			Send(c.conn, string(message))
		}
	}
}

func NewClient() net.Conn {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	go listen(conn)
	return conn
}

func listen(conn net.Conn) {
	for {
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}

func Send(conn net.Conn, text string) {
	fmt.Fprintf(conn, text + "\n")
}
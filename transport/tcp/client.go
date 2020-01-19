package tcp

import (
	"bufio"
	"net"
)

type Client struct {
	conn net.Conn
	send chan []byte
	receive chan[]byte
}

func NewClient() *Client {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
		receive: make(chan []byte, 256),
	}
	go client.listen()
	return client
}

func (c *Client) listen() {
	for {
		message, _ := bufio.NewReader(c.conn).ReadString('\n')
		c.receive <- []byte(message)
	}
}

func (c *Client) Send(msg []byte) {
	c.send <- msg
}
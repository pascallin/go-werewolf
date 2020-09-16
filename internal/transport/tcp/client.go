package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	conn    net.Conn
	send    chan []byte
	receive chan []byte
}

func NewClient() *Client {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	client := &Client{
		conn:    conn,
		send:    make(chan []byte, 256),
		receive: make(chan []byte, 256),
	}
	go client.listen()
	go client.pump()
	return client
}

func (c *Client) listen() {
	for {
		message, _ := bufio.NewReader(c.conn).ReadBytes('\n')
		c.receive <- message
	}
}

func (c *Client) Send(msg string) {
	// NOTE: need to add '\n' as line end
	c.send <- []byte(msg + "\n")
}

func (c *Client) pump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// TODO
				continue
			}
			_, err := c.conn.Write(message)
			if err != nil {
				log.Fatal(err)
			}
		case message, ok := <-c.receive:
			if !ok {
				// TODO
				continue
			}
			fmt.Println("client pump receive: " + string(message))
		}
	}
}

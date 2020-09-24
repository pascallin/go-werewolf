package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	conn    net.Conn
	send    chan []byte
	receive chan []byte
	writer	io.Writer
}

func NewClient(writer io.Writer) *Client {
	client := &Client{
		send:    make(chan []byte, 256),
		receive: make(chan []byte, 256),
	}
	return client
}

func (c *Client) Dia(url string) {
	// connect to this socket
	conn, _ := net.Dial("tcp", url)
	c.conn = conn
	go c.pump()
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
			fmt.Fprintf(c.writer, "client pump receive: " + string(message))
		}
	}
}

func (c *Client) Close() {
	c.conn.Close()
}
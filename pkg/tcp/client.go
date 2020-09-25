package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	conn    	net.Conn
	receive 	chan []byte
	writer		io.Writer
	CloseSignal	chan bool
}

func NewClient(writer io.Writer) *Client {
	client := &Client{
		receive: make(chan []byte),
		CloseSignal: make(chan bool),
		writer: writer,
	}
	return client
}

func (c *Client) Dia(url string) {

	// connect to this socket
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Fprintf(c.writer, "could not dial url: " + url + "\n")
		return
	}
	c.conn = conn
	defer c.Close()
	for {
		message, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(c.writer, "client pump receive: " + string(message))
	}
}

func (c *Client) Send(msg string) {
	// NOTE: need to add '\n' as line end
	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Client) Close() {
	c.conn.Close()
	c.CloseSignal <- true
}
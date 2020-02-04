package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type HubClient struct {
	hub  *Hub     // for server side hub chan change handle
	conn net.Conn // tcp conn
	send chan []byte
}

func (c *HubClient) Send(msg []byte) {
	c.send <- msg
}

// for server and client side
func (c *HubClient) writePump() {
	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// TODO
				return
			}
			fmt.Fprintf(c.conn, string(message))
		}
	}
}

func (c *HubClient) readPump() {
	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()
	for {
		message, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil || err == io.EOF {
			continue
		}
		c.hub.broadcast <- []byte(message)
	}
}

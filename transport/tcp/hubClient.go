package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type HubClient struct {
	hub *Hub // for server side hub chan change handle
	conn net.Conn // tcp conn
	send chan []byte
}

func (c *HubClient) Send(msg []byte) {
	c.send <- msg
}

// for server and client side
func (c *HubClient) writePump() {
	fmt.Println("HubClient writePump")
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
			fmt.Println("HubClient writePump get message: " + string(line))
			fmt.Fprintf(c.conn, string(message) + "\n")
		}
	}
}

func (c *HubClient) readPump() {
	fmt.Println("HubClient readPump")
	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()
	for {
		line, err := bufio.NewReader(c.conn).ReadBytes('\n')
		fmt.Println("HubClient readPump get message: " + string(line))
		if err != nil { return }
		c.hub.broadcast <- line
	}
}
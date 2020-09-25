package tcp

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net"
)

type ConnClient struct {
	ID		uuid.UUID
	conn 	net.Conn 		// tcp conn
}

func (c *ConnClient) send(message string) {
	fmt.Fprintf(c.conn, message)
}

func (c *ConnClient) close() {
	c.conn.Close()
}
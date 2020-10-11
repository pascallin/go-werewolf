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

func (c *ConnClient) send(message string) error {
	_, err := fmt.Fprintln(c.conn, message)
	return err
}

func (c *ConnClient) close() error {
	err := c.conn.Close()
	return err
}
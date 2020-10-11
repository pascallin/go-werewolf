package tcp

import (
	"os"
	"testing"
)

func TestClient_DiaAndSend(t *testing.T) {
	client := NewClient(os.Stdout)
	client.Dia("localhost:8080")

	client.Send("Hello")
}
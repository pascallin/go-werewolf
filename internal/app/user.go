package app

import (
	uuid "github.com/satori/go.uuid"

	"github.com/pascallin/go-wolvesgame/pkg/transport/tcp"
)

type User struct {
	Uid  uuid.UUID
	Nickname string
	CurrentRoom uuid.UUID
	TCPClient tcp.TCPClient
}

func NewUser(nickname string) User{
	return User{
		Uid: uuid.NewV4(),
		Nickname: nickname,
	}
}
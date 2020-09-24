package app

import (
	"github.com/pascallin/go-wolvesgame/pkg/tcp"
	uuid "github.com/satori/go.uuid"
)

type UserStatus int
const (
	Waiting UserStatus = iota
	Gaming
)
func (d UserStatus) String() string {
	return [...]string{"Waiting", "Gaming"}[d]
}

type User struct {
	Uid         uuid.UUID
	Nickname    string
	Status		UserStatus
	TCPClient   tcp.Client
}

func NewUser(nickname string) *User {
	return &User{
		Uid: uuid.NewV4(),
		Nickname: nickname,
		Status: Waiting,
	}
}

func (u *User) JoinGame() {
	u.Status = Gaming
}
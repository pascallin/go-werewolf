package player

import uuid "github.com/satori/go.uuid"

type Player struct {
	ID uuid.UUID
	Name string
}

func NewPlayer(id uuid.UUID, name string) Player {
	return Player{id, name}
}
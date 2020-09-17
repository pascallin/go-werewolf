package roles

import uuid "github.com/satori/go.uuid"

type Witch struct {
	Role
}

func (w *Witch) usePoison(playerID uuid.UUID) error {
	return nil
}

func (w *Witch) useAntidote(playerID uuid.UUID) error {
	return nil
}

func NewWitch() Role {
	var role = New("女巫(Witch)",  Good)
	return role
}
package game

type Status int

const (
	Waiting Status = iota	// wait for start
	Ready					// game started or restart
	Block					// waiting for player action
	Over					// game over
)

func (d Status) String() string {
	return [...]string{"Waiting", "Ready", "Block", "Over"}[d]
}
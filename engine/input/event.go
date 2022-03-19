package input

type Key int

const (
	UNKNOWN_KEY Key = iota
	LEFT
	RIGHT
	UP
	DOWN
	JUMP
	ESCAPE
)

type Action int

const (
	UNKNOWN_ACTION Action = iota
	PRESSED
	RELEASED
	REPEATED
)

type Event struct {
	Key    Key
	Action Action
}

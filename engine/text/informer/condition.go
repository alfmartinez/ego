package informer

type Condition = int

const (
	GAME_START_CONDITION Condition = iota
	CURRENT_ROOM_CONDITION
	DIRECTION_CONDITION
)

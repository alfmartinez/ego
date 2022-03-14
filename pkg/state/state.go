package state

type StateType int

const (
	StateIdle StateType = iota
	StateMove
)

type Updatable interface {
	Frame(x, y int)
}

type State interface {
	Update(Updatable) State
}

var states = make(map[StateType]func([]interface{}) State)

func RegisterStateFactory(code StateType, factory func([]interface{}) State) {
	states[code] = factory
}

func CreateState(code StateType, data ...interface{}) State {
	return states[code](data)
}

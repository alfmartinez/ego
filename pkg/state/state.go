package state

type StateType int

type HasStateType interface {
	Type() StateType
}

const (
	StateIdle StateType = iota
	StateMove
)

func (s StateType) Type() StateType {
	return s
}

type Updatable interface {
	Frame(x, y int)
}

type State interface {
	Update(Updatable) State
	HasStateType
}

var states = make(map[StateType]func([]interface{}) State)

func RegisterStateFactory(code StateType, factory func([]interface{}) State) {
	states[code] = factory
}

func CreateState(code StateType, data ...interface{}) State {
	return states[code](data)
}

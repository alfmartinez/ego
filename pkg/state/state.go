package state

type Updatable interface{}

type State interface {
	Update(Updatable) State
}

var states = make(map[string]func([]interface{}) State)

func RegisterStateFactory(name string, factory func([]interface{}) State) {
	states[name] = factory
}

func CreateState(name string, data ...interface{}) State {
	return states[name](data)
}

package text

import (
	"ego/engine/object"
	"ego/engine/observer"
	"ego/engine/state"
	"time"
)

func CreateLogic() object.GameObject {
	sm := state.CreateStateMachine()
	return &logic{sm}
}

type logic struct {
	state.StateMachine
}

func (l *logic) OnNotify(e observer.Event) {
	switch e.Type() {
	case observer.UPDATE:
		dt := e.Data()[0].(time.Duration)
		l.DoUpdate(l, dt)
	}
}

func (l *logic) Frame(x, y int) {

}

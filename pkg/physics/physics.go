package physics

import (
	"ego/pkg/context"
	"time"
)

type PhysicsEngine interface {
	Add(interface{})
	Advance(time.Duration)
}

func FromContext() PhysicsEngine {
	return context.GetContext().Get("physics").(PhysicsEngine)
}

func CreatePhysicsEngine() PhysicsEngine {
	return &phyiscsEngine{}
}

type phyiscsEngine struct {
	bodies []Body
}

func (e *phyiscsEngine) Add(i interface{}) {
	if body, ok := i.(Body); ok {
		e.bodies = append(e.bodies, body)
	}
}

func (e *phyiscsEngine) Advance(dt time.Duration) {

}

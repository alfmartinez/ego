package physics

import (
	"ego/pkg/context"
	"ego/pkg/movement"
	"fmt"
	"time"
)

const (
	NUM_STEPS    = 100
	SCALE_FACTOR = 1
)

type PhysicsEngine interface {
	Init()
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
	colliders    []Collider
	mobiles      map[int]Mobile
	matrices     map[int]M2
	countMobiles int
}

func (e *phyiscsEngine) Init() {
	e.colliders = make([]Collider, 0)
	e.mobiles = make(map[int]Mobile)
	e.countMobiles = 0
}

func (e *phyiscsEngine) Add(i interface{}) {
	if c, ok := i.(Collider); ok && !c.IsIgnored() {
		e.colliders = append(e.colliders, c)
	}
	if m, ok := i.(Mobile); ok {
		e.mobiles[e.countMobiles] = m
		e.countMobiles = e.countMobiles + 1
	}
}

func (e *phyiscsEngine) Advance(dt time.Duration) {
	e.initializeMatrices()
	if len(e.mobiles) == 0 {
		panic(fmt.Errorf("no body registered"))
	}
	step := dt.Seconds() / NUM_STEPS
	for s := 0; s < NUM_STEPS; s++ {
		for idx, matrix := range e.matrices {
			mobile := e.mobiles[idx]
			m := matrix.Advance(step)
			for cIndex, collider := range e.colliders {
				if idx != cIndex {
					collision := mobile.IsHit(collider)
					switch collision {
					case COLLISION_BOTTOM, COLLISION_TOP:
						m.S.Y = 0
					case COLLISION_LEFT, COLLISION_RIGHT:
						m.S.X = 0
					}
				}
			}
			e.matrices[idx] = m
		}
	}
	for idx, matrix := range e.matrices {
		body := e.mobiles[idx]
		body.SetPosition(movement.Location{
			X: matrix.P.X,
			Y: matrix.P.Y,
		})
		body.SetSpeed(movement.Speed{
			X: matrix.S.X,
			Y: matrix.S.Y,
		})
		body.SetAcceleration(movement.Acceleration{
			X: matrix.A.X,
			Y: matrix.A.Y,
		})
	}
}

func (e *phyiscsEngine) initializeMatrices() {
	matrices := make(map[int]M2, 0)
	for idx, mob := range e.mobiles {
		m := M2{}
		m.P = Vec2{mob.Position().X, mob.Position().Y}
		m.S = Vec2{mob.Speed().X, mob.Speed().Y}
		m.A = Vec2{mob.Acceleration().X, mob.Acceleration().Y}
		matrices[idx] = m
	}
	e.matrices = matrices
}

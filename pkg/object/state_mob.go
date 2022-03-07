package object

import (
	"ego/pkg/configuration"
	"ego/pkg/data"
	"ego/pkg/memory"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/sprite"
	"ego/pkg/state"
	"image"
)

type StateMob interface {
	GameObject
	memory.Memory
	data.Data
	movement.Movement
	motivator.NeedsCollection
}

type stateMob struct {
	state.StateMachine
	memory.Memory
	data.Data
	movement.Movement
	sprite.Sprite
	motivator.NeedsCollection
}

func CreateStateMob(config configuration.Mob) GameObject {
	name := config.Name
	position := image.Pt(config.Position.X, config.Position.Y)

	mobData := data.CreateData(name)
	mvmnt := movement.CreateMovement(position)
	memo := memory.CreateMemory()
	sprt := sprite.CreateSprite(config.Sprite)
	needs := motivator.CreateNeedsCollection()
	for _, need := range config.Needs {
		needs.AddNeed(motivator.CreateNeed(need.Type, need.Priority), need.Level)
	}
	sm := state.CreateStateMachine()

	return &stateMob{sm, memo, mobData, mvmnt, sprt, needs}
}

func (m *stateMob) Update() {
	m.StateMachine.Update(m)
}

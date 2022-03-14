package object

import (
	"ego/pkg/command"
	"ego/pkg/configuration"
	"ego/pkg/data"
	"ego/pkg/memory"
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/sprite"
	"ego/pkg/state"
	"image"
)

func init() {
	RegisterObjectFactory("Mob", CreateStateMob)
}

type StateMob interface {
	GameObject
	memory.Memory
	data.Data
	movement.Movement
	motivator.NeedsCollection
	sprite.Sprite
	command.CommandStream
}

type stateMob struct {
	state.StateMachine
	memory.Memory
	data.Data
	movement.Movement
	sprite.Sprite
	motivator.NeedsCollection
	command.CommandStream
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
		needs.AddNeed(motivator.CreateNeed(need.Type), need.Level)
	}
	sm := state.CreateStateMachine()
	stream := command.CreateCommandStream()

	return &stateMob{sm, memo, mobData, mvmnt, sprt, needs, stream}
}

func (m *stateMob) Update() {
	m.Execute()
	m.StateMachine.Update(m)
}

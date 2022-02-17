package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/state"
	"ego/pkg/utils"
)

type Mob struct {
	MobData
	memory.Memory
	state.StateMachine
}

func New(config configuration.Mob) *Mob {
	name := config.Name
	position := utils.Position{
		X: config.Position.X,
		Y: config.Position.Y,
	}
	mobData := MobData{name, position}
	stateMachine := state.CreateStateMachine()
	memory := memory.CreateMemory()
	return &Mob{
		mobData,
		memory,
		stateMachine,
	}
}

func (mob *Mob) GetName() string {
	return mob.name
}

func (m *Mob) Position() utils.Position {
	return m.position
}

func (mob *Mob) Render() {

}

package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/movement"
	"ego/pkg/mob/state"
	"ego/pkg/utils"
)

type Mob struct {
	*state.StateMachine
}

func New(config configuration.Mob) *Mob {
	name := config.Name
	position := utils.Position{
		X: config.Position.X,
		Y: config.Position.Y,
	}

	mobData := data.CreateMobData(name)
	movement := movement.CreateMovement(position)
	memory := memory.CreateMemory()
	stateMachine := state.CreateStateMachine(memory, mobData, movement)

	return &Mob{stateMachine}
}

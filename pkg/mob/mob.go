package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/state"
	"ego/pkg/utils"
)

type Mob struct {
	*data.Data
	*memory.Memory
	*state.StateMachine
}

func New(config configuration.Mob) *Mob {
	name := config.Name
	position := utils.Position{
		X: config.Position.X,
		Y: config.Position.Y,
	}

	mobData := data.CreateMobData(name, position)
	memory := memory.CreateMemory()
	stateMachine := state.CreateStateMachine(memory)

	return &Mob{
		mobData,
		memory,
		stateMachine,
	}
}

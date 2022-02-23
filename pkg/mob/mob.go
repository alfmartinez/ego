package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/motivator"
	"ego/pkg/mob/movement"
	"ego/pkg/mob/state"
	"ego/pkg/sprite"
	"ego/pkg/utils"
)

type Mob struct {
	state.StateMachine
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
	sprite := sprite.CreateSprite(config.Sprite)
	needs := motivator.CreateNeedsCollection()
	for _, need := range config.Needs {
		needs.AddNeed(motivator.CreateNeed(need.Type, need.Priority), need.Level)
	}
	stateMachine := state.CreateStateMachine(memory, mobData, movement, sprite, needs)

	return &Mob{stateMachine}
}

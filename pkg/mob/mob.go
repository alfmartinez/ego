package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/mob/data"
	"ego/pkg/mob/memory"
	"ego/pkg/mob/motivator"
	"ego/pkg/mob/movement"
	"ego/pkg/mob/state"
	"ego/pkg/renderer"
	"ego/pkg/sprite"
	"ego/pkg/terrain"
	"image"
)

type GameObject interface {
	Update(terrain.Terrain)
	Render(renderer.Renderer)
}

func CreateObject(config configuration.Mob) GameObject {
	name := config.Name
	position := image.Pt(config.Position.X, config.Position.Y)

	mobData := data.CreateMobData(name)
	movement := movement.CreateMovement(position)
	memory := memory.CreateMemory()
	sprite := sprite.CreateSprite(config.Sprite)
	needs := motivator.CreateNeedsCollection()
	for _, need := range config.Needs {
		needs.AddNeed(motivator.CreateNeed(need.Type, need.Priority), need.Level)
	}
	stateMachine := state.CreateStateMachine(memory, mobData, movement, sprite, needs)

	return stateMachine
}

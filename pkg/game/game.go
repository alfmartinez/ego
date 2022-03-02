package game

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type GameInterface interface {
	Start()
}

type GameObject interface {
	Update(terrain.Terrain)
	Render(renderer.Renderer)
}

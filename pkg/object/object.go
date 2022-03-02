package object

import (
	"ego/pkg/configuration"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type GameObject interface {
	Update(terrain.Terrain)
	Render(renderer.Renderer)
}

var objectFactories = make(map[string]func(configuration.Mob) GameObject)

func RegisterObjectFactory(name string, f func(configuration.Mob) GameObject) {
	objectFactories[name] = f
}

func CreateObject(config configuration.Mob) GameObject {
	name := config.Type
	return objectFactories[name](config)
}

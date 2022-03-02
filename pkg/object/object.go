package object

import (
	"ego/pkg/configuration"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type GameObject interface {
	Update(interface{}, terrain.Terrain)
	Render(interface{}, renderer.Renderer)
}

var objectFactories = make(map[string]func(configuration.Mob) GameObject)

func RegisterObjectFactory(name string, f func(configuration.Mob) GameObject) {
	objectFactories[name] = f
}

func CreateObject(config configuration.Mob) GameObject {
	name := config.Type
	return objectFactories[name](config)
}

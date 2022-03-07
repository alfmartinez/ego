package object

import (
	"ego/pkg/configuration"
)

type GameObject interface {
	Update()
}

var objectFactories = make(map[string]func(configuration.Mob) GameObject)

func RegisterObjectFactory(name string, f func(configuration.Mob) GameObject) {
	objectFactories[name] = f
}

func CreateObject(config configuration.Mob) GameObject {
	name := config.Type
	return objectFactories[name](config)
}

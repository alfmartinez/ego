package world

import "errors"

// World is the object defining where things exist and happen.

type World interface {
}

var worldFactories = make(map[string]func() World)

func RegisterWorld(name string, factory func() World) {
	worldFactories[name] = factory
}

func CreateWorld(name string) World {
	f, ok := worldFactories[name]
	if !ok {
		panic(errors.New("Unknown world " + name))
	}
	return f()
}

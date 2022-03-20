package object

import (
	"ego/engine/configuration"
	"ego/engine/observer"
	"fmt"
)

type GameObject interface {
	observer.Observer
}

var objectFactories = make(map[string]func(key string) GameObject)

func RegisterObjectFactory(name string, f func(key string) GameObject) {
	objectFactories[name] = f
}

func CreateObject(key string) GameObject {
	cfg := configuration.FromContext()
	name := cfg.GetString(fmt.Sprintf("mobs.%s.type", key))
	f, ok := objectFactories[name]
	if !ok {
		panic(fmt.Errorf("missing factory for object '%s'", name))
	}
	return f(key)
}

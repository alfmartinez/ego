package object

import (
	"ego/engine/context"
	"ego/engine/observer"
	"fmt"

	"github.com/spf13/viper"
)

type GameObject interface {
	observer.Observer
}

var objectFactories = make(map[string]func(key string) GameObject)

func RegisterObjectFactory(name string, f func(key string) GameObject) {
	objectFactories[name] = f
}

func CreateObject(key string) GameObject {
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	name := viper.GetString(fmt.Sprintf("mobs.%s.type", key))
	return objectFactories[name](key)
}

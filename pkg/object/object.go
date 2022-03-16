package object

import (
	"fmt"

	"github.com/spf13/viper"
)

type GameObject interface {
	Update()
}

var objectFactories = make(map[string]func(key string) GameObject)

func RegisterObjectFactory(name string, f func(key string) GameObject) {
	objectFactories[name] = f
}

func CreateObject(key string) GameObject {
	name := viper.GetString(fmt.Sprintf("mobs.%s.type", key))
	return objectFactories[name](key)
}

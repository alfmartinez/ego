package game

import (
	"ego/pkg/configuration"
)

func init() {
	RegisterGameFactory("file", func() (Game, error) {
		config, _ := configuration.ReadConfigurationFromFile("game.yml")
		return generateGame(config, CreateSampleGame), nil
	})
}

type Game interface {
	Start()
	Stop()
}

var factories = make(map[string]func() (Game, error))

func RegisterGameFactory(name string, f func() (Game, error)) {
	factories[name] = f
}

func CreateGame(name string) (Game, error) {
	return factories[name]()
}

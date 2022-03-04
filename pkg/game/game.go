package game

import (
	"ego/pkg/configuration"
)

func init() {
	RegisterGameFactory("file", func() Game {
		config, _ := configuration.ReadConfigurationFromFile("game.yml")
		return generateGame(config, CreateSampleGame)
	})
}

type Game interface {
	Start()
	Stop()
}

var factories = make(map[string]func() Game)

func RegisterGameFactory(name string, f func() Game) {
	factories[name] = f
}

func CreateGame(name string) Game {
	return factories[name]()
}

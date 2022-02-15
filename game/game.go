package game

import (
	"ego/configuration"
	"ego/mob"
)

type Game struct {
	Mobs []mob.Mob
}

func GenerateGame(config configuration.Configuration) *Game {

	var templates = make(map[string]*mob.Template)
	for _, x := range config.Templates {
		templates[x.Name] = mob.NewTemplate(x)
	}

	var mobs []mob.Mob
	for _, x := range config.Mobs {
		mob := mob.New(x)
		for _, template := range x.Templates {
			templates[template].Apply(mob)
		}
		mobs = append(mobs, *mob)
	}

	game := &Game{mobs}

	return game
}

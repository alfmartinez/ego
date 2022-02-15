package main

type Game struct {
	Mobs []Mob
}

func generateGame(config Configuration) *Game {

	var templates = make(map[string]*Template)
	for _, x := range config.Templates {
		templates[x.Name] = newTemplate(x)
	}

	var mobs []Mob
	for _, x := range config.Mobs {
		mob := newMob(x)
		for _, template := range x.Templates {
			templates[template].apply(mob)
		}
		mobs = append(mobs, *mob)
	}

	game := &Game{mobs}

	return game
}

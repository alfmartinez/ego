package game

func init() {
	RegisterGameFactory("viper", func() Game {
		return generateGame(CreateSampleGame)
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

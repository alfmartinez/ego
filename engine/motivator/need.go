package motivator

type Need int

const (
	None Need = iota
	Health
	Food
	Rest
	Learn
	Outdoor
	Indoor
	Recreation
	Beauty
)

var needs = map[string]Need{
	"health":     Health,
	"food":       Food,
	"rest":       Rest,
	"learn":      Learn,
	"outdoor":    Outdoor,
	"indoor":     Indoor,
	"recreation": Recreation,
	"beauty":     Beauty,
}

func CreateNeed(name string) Need {
	return needs[name]
}

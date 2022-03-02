package motivator

type Need interface {
	Name() string
	Priority() int
}

type need struct {
	name     string
	priority int
}

func CreateNeed(name string, priority int) Need {
	return &need{name, priority}
}

func (n *need) Name() string {
	return n.name
}

func (n *need) Priority() int {
	return n.priority
}

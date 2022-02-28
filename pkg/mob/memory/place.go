package memory

type PlaceMemory interface {
	IsExplored() bool
	Explore()
}

type placeMemory struct {
	explored int
}

func CreatePlaceMemory() PlaceMemory {
	return &placeMemory{explored: 0}
}

func (m *placeMemory) IsExplored() bool {
	return m.explored >= 10
}

func (m *placeMemory) Explore() {
	m.explored += 1
}

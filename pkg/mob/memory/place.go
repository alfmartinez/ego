package memory

type PlaceMemory struct {
	explored int
}

func (m *PlaceMemory) IsExplored() bool {
	return m.explored >= 10
}

func (m *PlaceMemory) Explore() {
	m.explored += 1
}

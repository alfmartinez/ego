package mob

import (
	"ego/pkg/utils"
)

type PlaceMemory struct {
	explored int
}

func (m *PlaceMemory) IsExplored() bool {
	return m.explored >= 10
}

func (m *PlaceMemory) Explore() {
	m.explored += 1
}

type Memory struct {
	places map[utils.Position]*PlaceMemory
}

func CreateMemory() *Memory {
	places := make(map[utils.Position]*PlaceMemory)
	return &Memory{places}
}

func (m *Memory) ExplorePlace(position utils.Position) {
	if place, ok := m.places[position]; ok {
		place.Explore()
	} else {
		m.places[position] = &PlaceMemory{explored: 1}
	}
}

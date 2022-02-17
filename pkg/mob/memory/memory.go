package memory

import (
	"ego/pkg/utils"
)

type Memory struct {
	places    map[utils.Position]*PlaceMemory
	interests []utils.Position
}

func CreateMemory() *Memory {
	places := make(map[utils.Position]*PlaceMemory)
	var interests []utils.Position
	return &Memory{places, interests}
}

func (m *Memory) UpdateInterests(position utils.Position, validate func(utils.Position) bool) {
	m.interests = position.Surrounding(5, validate)
}

func (m *Memory) ExplorePlace(position utils.Position) bool {
	if place, ok := m.places[position]; ok {
		place.Explore()
	} else {
		m.places[position] = &PlaceMemory{explored: 1}
	}
	return m.places[position].IsExplored()
}

func (m *Memory) SearchNextPositionToExplore() (utils.Position, bool) {
	for _, pos := range m.interests {
		if place, ok := m.places[pos]; !(ok && place.IsExplored()) {
			return pos, true
		}
	}

	return utils.Position{X: 0, Y: 0}, false
}

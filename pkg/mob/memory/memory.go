package memory

import (
	"ego/pkg/utils"
	"sort"
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
	m.interests = position.Surrounding(3, validate)
}

func (m *Memory) ExplorePlace(position utils.Position) bool {
	if place, ok := m.places[position]; ok {
		place.Explore()
	} else {
		m.places[position] = &PlaceMemory{explored: 1}
	}
	return m.places[position].IsExplored()
}

func (m *Memory) SearchNextPositionToExplore(position utils.Position) (utils.Position, bool) {
	sort.Slice(m.interests, func(i, j int) bool {
		first := m.interests[i]
		second := m.interests[j]
		if place, ok := m.places[second]; !(ok && place.IsExplored()) {
			return true
		}
		if place, ok := m.places[first]; !(ok && place.IsExplored()) {
			return false
		}
		return position.DistanceTo(first) < position.DistanceTo(second)
	})

	for _, pos := range m.interests {
		if place, ok := m.places[pos]; !(ok && place.IsExplored()) {
			return pos, true
		}
	}

	return utils.Position{X: 0, Y: 0}, false
}

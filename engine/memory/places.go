package memory

import (
	"ego/engine/movement"
)

type RemembersPlaces interface {
	HasExplored(movement.Positionnable) bool
	Explore(movement.Positionnable) bool
}

type places struct {
	places map[movement.Location]PlaceMemory
}

func CreatePlaces() RemembersPlaces {
	p := make(map[movement.Location]PlaceMemory)
	return &places{p}
}

func (m *places) HasExplored(location movement.Positionnable) bool {
	p, ok := m.places[location.Position()]
	return ok && p.IsExplored()
}

func (m *places) Explore(location movement.Positionnable) bool {
	place, ok := m.places[location.Position()]
	if !ok {
		place = CreatePlaceMemory()
		m.places[location.Position()] = place
	}
	place.Explore()
	return place.IsExplored()
}

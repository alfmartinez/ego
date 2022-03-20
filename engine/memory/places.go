package memory

import (
	"ego/engine/movement"
)

type RemembersPlaces interface {
	HasExplored(movement.Location) bool
	Explore(movement.Location) bool
}

type places struct {
	places map[movement.Location]PlaceMemory
}

func CreatePlaces() RemembersPlaces {
	p := make(map[movement.Location]PlaceMemory)
	return &places{p}
}

func (m *places) HasExplored(location movement.Location) bool {
	p, ok := m.places[location]
	return ok && p.IsExplored()
}

func (m *places) Explore(location movement.Location) bool {
	place, ok := m.places[location]
	if !ok {
		place = CreatePlaceMemory()
		m.places[location] = place
	}
	place.Explore()
	return place.IsExplored()
}

package mob

import (
	"ego/pkg/configuration"
	"ego/pkg/state"
	"ego/pkg/terrain"
	"ego/pkg/utils"
)

type Mob struct {
	name      string
	position  utils.Position
	memory    Memory
	state     state.State
	nextState state.State
}

func New(config configuration.Mob) *Mob {
	name := config.Name
	position := utils.Position{
		config.Position.X,
		config.Position.Y,
	}
	memory := CreateMemory()
	return &Mob{
		name,
		position,
		*memory,
		nil,
		nil,
	}
}

func (mob *Mob) GetName() string {
	return mob.name
}

func (mob *Mob) Update(grid terrain.Grid) {
	if mob.nextState != nil {
		mob.state = mob.nextState
		mob.nextState = nil
		mob.state.Enter()
	}
	if mob.state == nil {
		mob.nextState = state.CreateState("idle")
	} else {
		mob.nextState = mob.state.Update(mob, grid)
	}
}

func (mob *Mob) Render() {

}

// Actor Interface
func (m *Mob) Position() utils.Position {
	return m.position
}

func (m *Mob) HasFullyExplored(position utils.Position) bool {
	if !m.KnowsPlace(position) {
		return false
	}
	memory := m.GetPlaceMemory(position)
	if memory.explored < 100 {
		return false
	}
	return false
}

func (m *Mob) KnowsPlace(position utils.Position) bool {
	if _, ok := m.memory.places[position]; ok {
		return true
	}
	return false
}

func (m *Mob) GetPlaceMemory(position utils.Position) *PlaceMemory {
	return m.memory.places[position]
}

func (m *Mob) FindTileToExplore(g terrain.Grid) *terrain.Tile {
	return g.GetTile(m.position.Relative(1, 0))
}

func (m *Mob) ExecuteCommand(command string, options ...interface{}) {
	switch command {
	case "explore":
		position := options[0].(utils.Position)
		m.memory.ExplorePlace(position)
	}

}

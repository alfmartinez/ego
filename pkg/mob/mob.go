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
	state     state.State
	nextState state.State
}

func New(config configuration.Mob) *Mob {
	name := config.Name
	position := utils.Position{
		config.Position.X,
		config.Position.Y,
	}
	return &Mob{
		name,
		position,
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
	return true
}

func (m *Mob) FindTileToExplore(g terrain.Grid) *terrain.Tile {
	return nil
}

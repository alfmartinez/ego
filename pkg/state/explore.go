package state

import (
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type exploreState struct {
}

func (s *exploreState) Enter() {
	log.Print("Entering Explore State")
}

func (s *exploreState) Update(a Actor, g terrain.Grid) State {
	if a.HasFullyExplored(a.Position()) {
		tile := a.FindTileToExplore(g)
		if tile != nil {
			data := struct {
				destination utils.Position
			}{
				tile.Position,
			}
			return CreateState("move", data)
		}
		return CreateState("idle")
	}

	a.ExecuteCommand("explore", a.Position())

	return nil
}

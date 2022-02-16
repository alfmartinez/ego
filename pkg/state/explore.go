package state

import (
	"ego/pkg/terrain"
	"log"
)

type exploreState struct {
}

func (s *exploreState) Enter() {

}

func (s *exploreState) Update(a Actor, g terrain.Grid) State {
	log.Print("Explore state")
	return nil
}

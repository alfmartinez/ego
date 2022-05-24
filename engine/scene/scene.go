package scene

import (
	"github.com/alfmartinez/ego/engine/object"
)

type Scene struct {
	Label   string
	Objects []object.GameObject
}

func (s Scene) Start() {
	for _, o := range s.Objects {
		o.Start()
	}
}

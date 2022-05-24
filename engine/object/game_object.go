package object

import (
	"log"

	"github.com/alfmartinez/ego/engine/component"
)

type GameObject struct {
	Label    string
	Comps    []component.Component
	Children []GameObject
}

func (o GameObject) Start() {
	log.Printf("Starting object %s\n", o.Label)
	for _, c := range o.Children {
		c.Start()
	}
}

//go:generate go run generator.go

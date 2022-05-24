package object

import (
	"github.com/alfmartinez/ego/engine/component"
)

type GameObject struct {
	Label    string
	Comps    []component.Component
	Children []GameObject
}

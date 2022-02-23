package renderable

import "ego/pkg/utils"

type Renderable interface {
	Position() utils.Position
	Name() string
	Doing() string
	Path() string
}

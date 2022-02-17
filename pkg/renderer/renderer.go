package renderer

import "ego/pkg/utils"

type Renderable interface {
	Position() utils.Position
	Name() string
	Doing() string
}

type Renderer interface {
	Render(Renderable)
}

func CreateRenderer(name string) Renderer {
	renderers := map[string]func() Renderer{
		"null": func() Renderer {
			return &NullRenderer{}
		},
		"log": func() Renderer {
			return &LogRenderer{}
		},
	}
	return renderers[name]()
}

package renderer

import "ego/pkg/utils"

type Renderable interface {
	Position() utils.Position
	Name() string
	Doing() string
}

type Renderer interface {
	Init()
	Render(Renderable)
	Start()
}

func CreateRenderer(name string) Renderer {
	renderers := map[string]func() Renderer{
		"log": func() Renderer {
			return &LogRenderer{}
		},
		"fyne": func() Renderer {
			return &FyneRenderer{}
		},
	}
	return renderers[name]()
}

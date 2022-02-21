package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/utils"
)

type Renderable interface {
	Position() utils.Position
	Name() string
	Doing() string
}

type Renderer interface {
	Init()
	Render(Renderable)
	IsAsync() bool
	Start(chan bool)
	Refresh()
}

func CreateRenderer(config configuration.Renderer) Renderer {
	name := config.Type
	renderers := map[string]func(configuration.Renderer) Renderer{
		"log": func(config configuration.Renderer) Renderer {
			return &LogRenderer{}
		},
		"fyne": func(config configuration.Renderer) Renderer {
			display := display.CreateDisplay(config.Display)
			return &FyneRenderer{display: display}
		},
	}
	return renderers[name](config)
}

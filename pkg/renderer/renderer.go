package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/renderable"
)

type Renderer interface {
	Init()
	Render(renderable.Renderable)
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

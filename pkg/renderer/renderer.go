package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
)

type Renderer interface {
	Init()
	Render(renderable.Renderable)
	IsAsync() bool
	Start(chan bool)
	Refresh()
}

var rendererFactories = make(map[string]func(configuration.Renderer) Renderer)

func RegisterRendererFactory(name string, f func(configuration.Renderer) Renderer) {
	rendererFactories[name] = f
}

func CreateRenderer(config configuration.Renderer) Renderer {
	name := config.Type
	return rendererFactories[name](config)
}

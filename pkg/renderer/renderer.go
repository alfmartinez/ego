package renderer

import (
	"ego/pkg/configuration"
)

type Renderer interface {
	Init()
	IsAsync() bool
	Start(chan bool)
	Render(RenderTree)
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

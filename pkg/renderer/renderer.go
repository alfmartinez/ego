package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/render"
	"errors"
	"log"
)

type Renderer interface {
	Init()
	IsAsync() bool
	Start(chan bool)
	Render(render.RenderTree)
	Refresh()
	Close()
}

var rendererFactories = make(map[string]func(configuration.Renderer) Renderer)

func RegisterRendererFactory(name string, f func(configuration.Renderer) Renderer) {
	rendererFactories[name] = f
}

func CreateRenderer(config configuration.Renderer) Renderer {
	name := config.Type
	f, ok := rendererFactories[name]
	if !ok {
		log.Printf("%+v", rendererFactories)
		panic(errors.New("Can find factory " + name))
	}
	return f(config)
}

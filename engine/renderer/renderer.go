package renderer

import (
	"ego/engine/context"
	"errors"
)

type Renderer interface {
	Init()
	Start()
	Render(interface{})
	Refresh()
	Close()
}

var rendererFactories = make(map[string]func() Renderer)

var currentRenderer Renderer

func GetRenderer() Renderer {
	return currentRenderer
}

func RegisterRendererFactory(name string, f func() Renderer) {
	rendererFactories[name] = f
}

func FromContext() Renderer {
	return context.GetContext().Get("renderer").(Renderer)
}

func CreateRenderer(name string) Renderer {
	f, ok := rendererFactories[name]
	if !ok {
		panic(errors.New("Cannot find renderer factory " + name))
	}
	currentRenderer := f()
	return currentRenderer
}

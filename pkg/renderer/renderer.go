package renderer

import (
	"ego/pkg/display"
	"errors"
)

type Renderer interface {
	Init()
	Start(chan bool)
	Render(display.Displayable)
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

func CreateRenderer(name string) Renderer {
	f, ok := rendererFactories[name]
	if !ok {
		panic(errors.New("Cannot find renderer factory " + name))
	}
	currentRenderer := f()
	return currentRenderer
}

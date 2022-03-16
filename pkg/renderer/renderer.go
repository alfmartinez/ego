package renderer

import (
	"ego/pkg/render"
	"errors"

	"github.com/spf13/viper"
)

type Renderer interface {
	Init()
	Start(chan bool)
	Render(render.RenderTree)
	Refresh()
	Close()
}

var rendererFactories = make(map[string]func() Renderer)

func RegisterRendererFactory(name string, f func() Renderer) {
	rendererFactories[name] = f
}

func CreateRenderer() Renderer {
	name := viper.GetString("renderer.type")
	f, ok := rendererFactories[name]
	if !ok {
		panic(errors.New("Cannot find renderer factory " + name))
	}
	return f()
}

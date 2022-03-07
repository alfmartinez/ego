package renderable

import "ego/pkg/renderer"

type Renderable interface {
	Render(interface{}, renderer.Renderer)
}

var factories = make(map[string]func() Renderable)

func RegisterRenderableFactory(name string, f func() Renderable) {
	factories[name] = f
}

func CreateRenderable(name string) Renderable {
	return factories[name]()
}

package renderable

type Renderable interface {
	Render()
}

var factories = make(map[string]func() Renderable)

func RegisterRenderableFactory(name string, f func() Renderable) {
	factories[name] = f
}

func CreateRenderable(name string) Renderable {
	return factories[name]()
}

package renderer

type RenderTree interface {
	Apply(func(interface{}))
}

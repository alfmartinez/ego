package renderer

type Renderable interface {
}

type Renderer interface {
	Render(Renderable)
}

func CreateRenderer(name string) Renderer {
	renderers := map[string]func() Renderer{
		"null": func() Renderer {
			return &NullRenderer{}
		},
	}
	return renderers[name]()
}

package console

import "ego/engine/renderer"

func Register() {
	renderer.RegisterRendererFactory("console", func() renderer.Renderer {
		return &console{}
	})
}

type console struct {
}

// Start implements renderer.Renderer
func (*console) Start() {

}

// Close implements renderer.Renderer
func (*console) Close() {

}

// Init implements renderer.Renderer
func (*console) Init() {

}

// Render implements renderer.Renderer
func (*console) Render(interface{}) {

}

// Refresh implements renderer.Renderer
func (*console) Refresh() {

}

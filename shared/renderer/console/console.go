package console

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/renderer"
)

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
func (r *console) Init() {

}

// Render implements renderer.Renderer
func (r *console) Render(s interface{}) {
	fmt.Print(s.(string))
}

// Refresh implements renderer.Renderer
func (r *console) Refresh() {

}

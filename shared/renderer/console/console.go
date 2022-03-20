package console

import (
	"ego/engine/renderer"
	"fmt"
	"strings"
)

func Register() {
	renderer.RegisterRendererFactory("console", func() renderer.Renderer {
		return &console{}
	})
}

type console struct {
	buffer []string
}

// Start implements renderer.Renderer
func (*console) Start() {

}

// Close implements renderer.Renderer
func (*console) Close() {

}

// Init implements renderer.Renderer
func (r *console) Init() {
	r.buffer = make([]string, 0)
}

// Render implements renderer.Renderer
func (r *console) Render(s interface{}) {
	r.buffer = append(r.buffer, s.(string))
}

// Refresh implements renderer.Renderer
func (r *console) Refresh() {
	content := strings.Join(r.buffer, "\n")
	fmt.Print(content)
	r.buffer = make([]string, 0)
}

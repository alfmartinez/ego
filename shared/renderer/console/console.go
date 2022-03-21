package console

import (
	"ego/engine/configuration"
	"ego/engine/renderer"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Register() {
	renderer.RegisterRendererFactory("console", func() renderer.Renderer {
		langKey := configuration.FromContext().GetString("renderer.lang")
		return &console{language.Make(langKey)}
	})
}

type console struct {
	lang language.Tag
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
	p := message.NewPrinter(r.lang)
	p.Print(s.(string))
}

// Refresh implements renderer.Renderer
func (r *console) Refresh() {

}

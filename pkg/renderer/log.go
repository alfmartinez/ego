package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"log"
)

func init() {
	RegisterRendererFactory("log", func(config configuration.Renderer) Renderer {
		return &LogRenderer{}
	})
}

type LogRenderer struct {
}

func (r *LogRenderer) IsAsync() bool {
	return false
}

func (r *LogRenderer) Init() {

}

func (r *LogRenderer) Start(exit chan bool) {

}

func (r *LogRenderer) Refresh() {
}

func (r *LogRenderer) Render(s renderable.Renderable) {
	if s.Name() != "" {
		log.Printf("- %s %s at %+v", s.Name(), s.Doing(), s.Position())
	}
}

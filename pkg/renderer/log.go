package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/utils"
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

func (r *LogRenderer) Init() {}

func (r *LogRenderer) Start(exit chan bool) {}

func (r *LogRenderer) Refresh() {}

func (r *LogRenderer) Render(tree RenderTree) {
	tree.Apply(func(s utils.Tree) {
		log.Printf("%+v", s)
	})

}

package renderer

import "log"

type LogRenderer struct {
}

func (r *LogRenderer) IsAsync() bool {
	return false
}

func (r *LogRenderer) Init() {

}

func (r *LogRenderer) Start() {

}

func (r *LogRenderer) Refresh() {
}

func (r *LogRenderer) Render(s Renderable) {
	log.Printf("- %s %s at %+v", s.Name(), s.Doing(), s.Position())
}

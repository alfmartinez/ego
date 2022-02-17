package renderer

import "log"

type LogRenderer struct {
}

func (r *LogRenderer) Render(s Renderable) {
	log.Printf("- %s %s at %+v", s.Name(), s.Doing(), s.Position())
}

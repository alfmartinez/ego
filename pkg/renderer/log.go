package renderer

import "log"

type LogRenderer struct {
}

func (r *LogRenderer) Render(s Renderable) {
	log.Printf("Mob %s Doing %s at %+v", s.Name(), s.Doing(), s.Position())
}

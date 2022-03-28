package text

import (
	"github.com/alfmartinez/ego/engine/observer"
	"github.com/alfmartinez/ego/engine/renderer"
	"time"
)

type textGame struct {
	subject  observer.Subject
	stopLoop bool
}

func (g *textGame) Start() {
	g.subject.Register(g)
	g.loop()
}

func (g *textGame) Stop() {

}

func (g *textGame) OnNotify(e observer.Event) {
	switch e.Type() {
	case observer.EXIT:
		g.stopLoop = true
	case observer.RENDER:
		renderer.FromContext().Refresh()
	}
}

func (g *textGame) loop() {
	var lastUpdate = time.Now()
	for !g.stopLoop {
		g.subject.NotifyAll(observer.CreateEvent(observer.UPDATE, time.Since(lastUpdate)))
		lastUpdate = time.Now()
		g.subject.NotifyAll(observer.CreateEvent(observer.RENDER))
	}
}

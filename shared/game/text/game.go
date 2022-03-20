package text

import (
	"ego/engine/observer"
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
	}
}

func (g *textGame) loop() {
	var lastUpdate = time.Now()
	for !g.stopLoop {
		g.subject.NotifyAll(observer.CreateEvent(observer.UPDATE, time.Since(lastUpdate)))
		lastUpdate = time.Now()
	}
}

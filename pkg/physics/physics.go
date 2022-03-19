package physics

import (
	"time"
)

type Engine interface {
	Init()
	Add(interface{})
	Advance(time.Duration)
}

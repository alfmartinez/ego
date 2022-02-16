package behaviour

import (
	"errors"
)

type Behaviour interface {
	GetName() string
	Evaluate()
	IsOver() bool
	Reset()
	Next() string
}

type behaviour struct {
	Name string
	data struct {
		duration int
		timerOn  bool
	}
}

func (b *behaviour) startTimer(duration int) {
	b.data.duration = duration
	b.data.timerOn = true
}

func (b *behaviour) decrementTimer() {
	b.data.duration--
}

func (b *behaviour) isTimerOver() bool {
	return b.data.duration == 0
}

func (b *behaviour) isTimerRunning() bool {
	return b.data.timerOn
}

type notimplemented struct {
	behaviour
}

func (behaviour *notimplemented) GetName() string {
	return behaviour.Name
}

func (behaviour *notimplemented) Evaluate() {}

func (behaviour *notimplemented) Reset() {}

func (behaviour *notimplemented) IsOver() bool {
	return false
}

func (behaviour *notimplemented) Next() string {
	return "idle"
}

var behaviours map[string]func(s string) Behaviour

func New(name string) Behaviour {
	if behaviours == nil {
		behaviours = map[string]func(s string) Behaviour{
			"idle": func(s string) Behaviour {
				return &idle{behaviour{
					Name: s,
				}}
			},
			"wait": func(s string) Behaviour {
				return &wait{behaviour{
					Name: s,
				}}
			},
			"sleep": func(s string) Behaviour {
				return &sleep{behaviour{
					Name: s,
				}}
			},

			"eat": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"attack": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"defend": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"guard": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"cast": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"shield": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"guess": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
			"bark": func(s string) Behaviour {
				return &notimplemented{behaviour{
					Name: s,
				}}
			},
		}
	}

	if (behaviours[name]) == nil {
		panic(errors.New("Behaviour unknown " + name))
	}

	return behaviours[name](name)
}

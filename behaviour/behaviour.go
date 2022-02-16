package behaviour

import (
	"errors"
)

type Behaviour interface {
	GetName() string
	Evaluate() string
}

type behaviour struct {
	Name string
	data struct {
		duration int
	}
}

func (b *behaviour) SetDuration(duration int) {
	b.data.duration = duration
}

func (b *behaviour) decrementDuration() int {
	b.data.duration--
	return b.data.duration
}

type notimplemented struct {
	behaviour
}

func (behaviour *notimplemented) GetName() string {
	return behaviour.Name
}

func (behaviour *notimplemented) Evaluate() string {
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

package behaviour

import (
	"errors"
)

type Behaviour interface{}

type behaviour struct {
	Name string
}

type idle behaviour
type wait behaviour
type sleep behaviour
type notimplemented behaviour

var behaviours map[string]func(s string) Behaviour

func New(name string) Behaviour {
	if behaviours == nil {
		behaviours = map[string]func(s string) Behaviour{
			"idle":  func(s string) Behaviour { return &idle{s} },
			"wait":  func(s string) Behaviour { return &wait{s} },
			"sleep": func(s string) Behaviour { return &sleep{s} },

			"eat":    func(s string) Behaviour { return &notimplemented{s} },
			"attack": func(s string) Behaviour { return &notimplemented{s} },
			"defend": func(s string) Behaviour { return &notimplemented{s} },
			"guard":  func(s string) Behaviour { return &notimplemented{s} },
			"cast":   func(s string) Behaviour { return &notimplemented{s} },
			"shield": func(s string) Behaviour { return &notimplemented{s} },
			"guess":  func(s string) Behaviour { return &notimplemented{s} },
			"bark":   func(s string) Behaviour { return &notimplemented{s} },
		}
	}

	if (behaviours[name]) == nil {
		panic(errors.New("Behaviour unknown " + name))
	}

	return behaviours[name](name)
}

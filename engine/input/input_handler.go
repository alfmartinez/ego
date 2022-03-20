package input

import "fmt"

type InputHandler interface {
	Handle(Event)
}

var handlerFactories = make(map[string]func() InputHandler)

func RegisterInputHandler(name string, f func() InputHandler) {
	handlerFactories[name] = f
}

func CreateInputHandler(name string) InputHandler {
	f, ok := handlerFactories[name]
	if !ok {
		panic(fmt.Errorf("Unknown input handler %s", name))
	}
	return f()
}

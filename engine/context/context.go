package context

import (
	"fmt"
	"sync"
)

var (
	mu             sync.Mutex
	contexts       = make(map[string]Context)
	currentContext Context
)

func Set(name string, value interface{}) {
	GetContext().Set(name, value)
}

func Get(name string) interface{} {
	return GetContext().Get(name)
}

type Context interface {
	Set(string, interface{})
	Get(string) interface{}
}

func CreateAndRegisterContext(name string) {
	ctx := CreateContext()
	RegisterContext(name, ctx)
}

func CreateContext() Context {
	values := make(map[string]interface{})
	return &context{values}
}

func RegisterContext(name string, ctx Context) {
	mu.Lock()
	contexts[name] = ctx
	currentContext = ctx
	mu.Unlock()
	fmt.Printf("Current context %#v", currentContext)
}

func SwitchContext(name string) {
	mu.Lock()
	ctx := contexts[name]
	currentContext = ctx
	mu.Unlock()
}

func GetContext() Context {
	if currentContext == nil {
		panic(fmt.Errorf("Current context is not set !"))
	}
	return currentContext
}

type context struct {
	values map[string]interface{}
}

func (c *context) Set(name string, value interface{}) {
	c.values[name] = value
}

func (c *context) Get(name string) interface{} {
	comp, ok := c.values[name]
	if !ok {
		panic(fmt.Errorf("Missing component in context for key '%s'", name))
	}
	return comp
}

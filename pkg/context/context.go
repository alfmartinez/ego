package context

func Set(name string, value interface{}) {
	GetContext().Set(name, value)
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

var contexts = make(map[string]Context)
var currentContext string

func RegisterContext(name string, ctx Context) {
	contexts[name] = ctx
	currentContext = name
}

func GetContext() Context {
	return contexts[currentContext]
}

type context struct {
	values map[string]interface{}
}

func (c *context) Set(name string, value interface{}) {
	c.values[name] = value
}

func (c *context) Get(name string) interface{} {
	return c.values[name]
}

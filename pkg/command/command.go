package command

type Command interface {
	Execute() bool
}

func CreateCommand(e func() bool) Command {
	return &command{e}
}

type command struct {
	f func() bool
}

func (c *command) Execute() bool {
	return c.f()
}

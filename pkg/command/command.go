package command

type Command interface {
	Execute() bool
	Abort()
}

func CreateCommand(e func() bool) Command {
	return &command{e, false}
}

type command struct {
	f       func() bool
	aborted bool
}

func (c *command) Abort() {
	c.aborted = true
}

func (c *command) Execute() bool {
	if c.aborted {
		return true
	}
	return c.f()
}

package text

type (
	Command struct {
		Direction *Direction `@@`
	}
)

func (c *Command) String() string {
	return c.Direction.Value
}

func ParseCommand(cmd string) *Command {
	ast := &Command{}
	parser := BuildCommandParser()
	err := parser.ParseString("", cmd, ast)
	if err != nil {
		panic(err)
	}
	return ast
}

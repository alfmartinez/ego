package grammar

type (
	Command struct {
		Verb   string `@Ident`
		Target string `@Ident?`
	}
)

func (c *Command) String() string {
	return c.Verb + " " + c.Target

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

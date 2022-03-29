package grammar

import (
	"fmt"
	"strings"
)

type (
	Command struct {
		Direction *Direction `  @@`
		Action    *Action    `| @@`
	}
	Action struct {
		Verb   string      `@Verb`
		Target *Designator `@@`
	}
)

func (c *Command) String() string {
	if c.Direction != nil {
		return c.Direction.Value
	}
	if c.Action != nil {
		return c.Action.Verb + " " + strings.Join(c.Action.Target.Elements, " ")
	}
	panic(fmt.Errorf("Can't display command"))
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

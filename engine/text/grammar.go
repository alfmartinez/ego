package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Story struct {
	Title string `@String`
	Rooms []Room `@@*`
}

type Room struct {
	Ident string `@Room`
	Props []Prop `@@*`
}

type Prop struct {
	Name string `@Prop`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Room", `(.+) est une pièce.`, nil},
			{"Prop", `(.+) est là.`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"nl", `\n`, nil},
			{"whitespace", `\s+`, nil},
		},
	})
)

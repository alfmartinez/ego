package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Story struct {
	Title string `@Title`
	Rooms []Room `@@*`
}

type Room struct {
	Name        string `@Room`
	Description string `@Description`
	Props       []Prop `@@*`
}

type Prop struct {
	Name string `@Prop`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			lexer.Include("Discard"),
			{"Title", `"(\\"|[^"])*"`, nil},
			{"Room", `(.*) est une pièce.`, lexer.Push("Room")},
		},
		"Room": {
			lexer.Include("Discard"),
			{"Ident", `\1`, nil},
			{"Description", `"(\\"|[^"])*"`, nil},
			{"Prop", `(.*) est là.`, lexer.Push("Prop")},
		},
		"Prop": {
			lexer.Include("Discard"),
			{"Ident", `\1`, nil},
		},
		"Discard": {
			{"nl", `\n`, nil},
			{"whitespace", `\s+`, nil},
		},
	})
)

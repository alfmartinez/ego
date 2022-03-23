package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Title       string       `@String`
	Definitions []Definition `@@*`
}

type Definition struct {
	Subject    []string `@Ident* "est":Verb`
	Complement []string `@Ident* "."`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `\/\/.+\n`, nil},
			{"newLine", `\n+`, nil},
			{"whitespace", `\s+`, nil},
			{"Punct", `[\.]`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"Verb", `(est)`, nil},
			{"Ident", `[\p{L}]+`, nil},
		},
	})
)

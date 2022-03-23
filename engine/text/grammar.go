package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Grammar struct {
	Title string `@String`
}

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"comment", `\/\/.+\n`, nil},
			{"punct", `[\.\,\s]+`, nil},
			{"whitespace", `\s+`, nil},
			{"nl", `\n`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
			{"Ident", `([\p{L}]+\s*)+`, nil},
		},
	})
)

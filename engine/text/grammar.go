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
			{"eol", `\n`, nil},
			{"whitespace", `\s+`, nil},
			{"String", `"(\\"|[^"])*"`, nil},
		},
	})
)

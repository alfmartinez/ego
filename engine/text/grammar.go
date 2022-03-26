package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type (
	// Define a Grammar for parsing
	Grammar struct {
		//
		World World `@@ EOF`
	}

	World struct {
		Statements []*Statement `(@@ EOL)*`
	}

	Statement struct {
		Title    string    `(  @String`
		Room     *Room     ` | @@ `
		Item     *Item     ` | @@ `
		Describe *Describe ` | @@ )`
	}

	Room struct {
		Designator  *Designator `@@ Is Room`
		Description string      `@String?`
	}

	Item struct {
		Designator  *Designator `@@ Is Here`
		Description string      `@String?`
	}

	Describe struct {
		Designator  *Designator `Describe @@ `
		Description string      `Is @String`
	}

	Designator struct {
		KeyWords []string `@Ident+`
	}
)

var (
	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"Comment", `//[^\n]*\n`, nil},
			{"String", `"[^"]*"`, nil},
			{"Is", `est`, nil},
			//	{"Article", `((U|u)n(e)?|(L|l)(a|e))`, nil},
			{"Room", `un lieu\.`, nil},
			{"Here", `ici\.`, nil},
			{"Describe", `La description`, nil},
			{"Ident", `\p{L}+`, nil},
			{"Punct", `[\.]+`, nil},
			{"Whitespace", `[ \t]+`, nil},
			{"EOL", `\n+`, nil},
		},
	})
)

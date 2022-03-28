package text

import (
	"github.com/alecthomas/participle/v2/lexer"
)

// Defines a grammar for natural text parsing.
// 1. 	S = DP + VP 			: The Chamber is a room.
// 2. 	VP = Verb + DP			: is a room
// 3a. 	DP = DP + which + VP	: The cart is a vehicle which is movable.
// 3b. 	DP = DP + who + VP		: The Janitor is a male person who has the blue key.
// 4. 	DP = DP + RP 			: An open container in the ballroom
// 5a. 	RP = Preposition + DP 	: in a container
// 5b. 	RP = Participle + DP 	: worn by Mr Darcy
// 6 	S = RP + Verb + DP 		: In the Garden is a sunflower

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

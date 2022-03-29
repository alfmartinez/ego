package text

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
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
		Sentence *Sentence `@@`
		Title    string    `| @String`
		Section  *Section  `| @@`
	}

	Section struct {
		Number int    `Section @Number`
		Title  string `"-" (@String | @Ident)`
		Tag    string `("-" (@String | @Ident*))? `
	}

	Sentence struct {
		DP          *DescriptionPhrase `@@`
		VP          *VerbPhrase        `@@ "."`
		Description string             `@String?`
	}

	DescriptionPhrase struct {
		Designator *Designator     ` @@ `
		List       *List           ` @@?`
		Complex    *ComplexPhrase  ` @@?`
		Relation   *RelativePhrase ` @@?`
	}

	List struct {
		Elements []*Designator `("," @@)*`
		Last     *Designator   `"and" @@`
	}

	Designator struct {
		Elements []string `{ @Ident }`
	}

	ComplexPhrase struct {
		VP *VerbPhrase `Determiner @@`
	}

	RelativePhrase struct {
		Relation string      `@Relation`
		Related  *Designator `@@`
	}

	VerbPhrase struct {
		Verb string             `@Verb`
		DP   *DescriptionPhrase `@@`
	}
)

var (
	verbs       = []string{"is", "has", "carries"}
	articles    = []string{"a", "an", "the", "The", "An", "A"}
	determiners = []string{"which", "who"}
	relations   = []string{"of", "in"}

	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"[^"]*"`, nil},
			{"Separator", `(\,|and\b)`, nil},
			{"Section", `Section\b`, nil},
			{"Relation", "(" + strings.Join(relations, "|") + `)\b`, nil},
			{"Determiner", "(" + strings.Join(determiners, "|") + `)\b`, nil},
			{"Article", "(" + strings.Join(articles, "|") + `)\b`, nil},
			{"Verb", "(" + strings.Join(verbs, "|") + `)\b`, nil},
			{"Ident", `[\p{L}]+`, nil},
			{"Number", `[0-9]+`, nil},
			{"Punct", `[\.\-]`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `[ \t]+`, nil},
		},
	})
)

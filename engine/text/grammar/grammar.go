package grammar

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
		Title           string           `@String`
		ValueDefinition *ValueDefinition `| @@`
		KindDefinition  *KindDefinition  `| @@`
		Direction       *Connector       `| @@`
		Sentence        *Sentence        `| @@`
		Section         *Section         `| @@`
		Test            *Test            `| @@`
		Description     *Description     `| @@`
	}

	ValueDefinition struct {
		Name *Designator `@@ "is" Kind "of" "value"`
		With []string    `("with" "values" (@Ident Separator?)+)? "."`
	}

	KindDefinition struct {
		Name   *Designator `@@ "is" Kind "of"`
		Parent *Designator `@@`
		With   []*Property `("with" (@@ Separator?)*)? "."`
	}

	Property struct {
		Property *Designator `@@`
		Value    string      `@String`
	}

	Description struct {
		Target      *Designator `"description" "of" @@`
		Description string      `"is" @String "."?`
	}

	Connector struct {
		Direction   *Direction  `@@`
		Origin      *Designator `"of" @@ "is"`
		Target      *Designator `@@ "."`
		Description string      `@String?`
	}

	Direction struct {
		Value string `@Direction`
	}

	Test struct {
		Commands []string `Test {@String}`
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
	verbs       = []string{"is", "has", "carries", "look"}
	articles    = []string{"a", "an", "the", "The", "An", "A"}
	determiners = []string{"which", "who"}
	relations   = []string{"of", "in", "with"}
	directions  = []string{
		"north", "south", "east", "west",
		"North", "South", "East", "West",
	}

	def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"String", `"[^"]*"`, nil},
			{"Separator", `(\,|and\b)`, nil},
			{"Section", `Section\b`, nil},
			{"Test", `Test me with`, nil},
			{"Kind", `kind\b`, nil},
			{"Direction", "(" + strings.Join(directions, "|") + `)\b`, nil},
			{"Relation", "(" + strings.Join(relations, "|") + `)\b`, nil},
			{"Determiner", "(" + strings.Join(determiners, "|") + `)\b`, nil},
			{"Article", "(" + strings.Join(articles, "|") + `)\b`, nil},
			{"Verb", "(" + strings.Join(verbs, "|") + `)\b`, nil},
			{"Ident", `[\p{L}]+`, nil},
			{"Number", `[0-9]+`, nil},
			{"Punct", `[\.\-\(\)\:]`, nil},
			{"EOL", `\n+`, nil},
			{"Whitespace", `[ \t]+`, nil},
		},
	})
)

func (d *Designator) Get() string {
	return strings.Join(d.Elements, " ")
}

func (d *Direction) Direct() string {
	return d.Value
}

func (d *Direction) Reverse() string {
	return DirectionReverse[d.Value]
}

var DirectionReverse = map[string]string{
	"east":      "west",
	"south":     "north",
	"west":      "east",
	"north":     "south",
	"up":        "down",
	"down":      "up",
	"northeast": "southwest",
	"northwest": "southeast",
	"southwest": "northeast",
	"southeast": "northwest",
}

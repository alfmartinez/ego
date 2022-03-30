// Text package defines a set of tools to create states for state machine from a text file
// Synopsis :
// - Define grammar
// - Parse text file(s)
// - Generate States
package text

import (
	"github.com/alfmartinez/ego/engine/text/grammar"
	"github.com/alfmartinez/ego/engine/text/informer"
	"io"
)

type Story interface {
	Start()
	Test()
	SetWriter(io.Writer)
}

func CreateStory(filepath string, debug bool) Story {
	ast := grammar.ParseFile(filepath)
	story := informer.CreateRuleSemantix(debug).BuildStory(ast)
	return story
}

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
	"os"
)

type Story interface {
	Start()
	Test()
	SetWriter(io.Writer)
}

func CreateStory(filepaths []string, debug bool, tokens bool) Story {
	analyzer := informer.CreateRuleSemantix(debug)

	for _, filepath := range filepaths {
		if tokens {
			f, _ := os.Open(filepath)
			grammar.ParseTokens(f)
		}
		ast := grammar.ParseFile(filepath)
		analyzer.BuildStory(ast)
	}

	story := analyzer.GetStory()
	return story
}

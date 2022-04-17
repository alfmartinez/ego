// Text package defines a set of tools to create states for state machine from a text file
// Synopsis :
// - Define grammar
// - Parse text file(s)
// - Generate States
package text

import (
	"ego/engine/text/informer"
	"ego/engine/text/parser"
	"io"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type informerListener struct {
	parser.BaseInformerListener
}

type Story interface {
	Start()
	SetWriter(io.Writer)
	Test()
}

func CreateStory(filepaths []string, debug bool, tokens bool) Story {
	l := &informerListener{}
	for _, filepath := range filepaths {
		is, _ := antlr.NewFileStream(filepath)
		// Create the Lexer
		lexer := parser.NewInformerLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		// Create the Parser
		p := parser.NewInformerParser(stream)

		// Finally parse the expression
		antlr.ParseTreeWalkerDefault.Walk(l, p.Start())

	}

	sem := informer.CreateRuleSemantix(true)

	return sem.GetStory()
}

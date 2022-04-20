package text

import (
	"ego/engine/text/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type informerVisitor struct {
	parser.BaseInformerVisitor
	rulebooks []string
}

func (v *informerVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic("implement visit children")
}
func (v *informerVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("implement visit terminal")
}
func (v *informerVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement visit error node")
}

func (v *informerVisitor) VisitRulebook(ctx *parser.RulebookContext) interface{} {
	key := ctx.Identifier().GetText()
	v.rulebooks = append(v.rulebooks, key)
	fmt.Printf("Rulebook %s created \n", key)
	return nil
}

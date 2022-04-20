package text

import (
	"ego/engine/text/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type informerVisitor struct {
	parser.BaseInformerVisitor
	rulebooks []string
}

func (v *informerVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *informerVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *informerVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, c := range node.GetChildren() {
		switch t := c.(type) {
		case *parser.StatementContext,
			*parser.RulebookContext,
			*parser.ActivityContext,
			*antlr.TerminalNodeImpl:
			v.Visit(c.(antlr.ParseTree))
		default:
			panic(fmt.Errorf("unknown type %T", t))
		}
	}
	return nil
}

func (v *informerVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *informerVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}
func (v *informerVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement visit error node")
}

func (v *informerVisitor) VisitRulebook(ctx *parser.RulebookContext) interface{} {
	key := strings.Join(v.Visit(ctx.Identifier()).([]string), " ")
	v.rulebooks = append(v.rulebooks, key)
	fmt.Printf("Rulebook %s created \n", key)
	return nil
}

func (v *informerVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	w := ctx.WORD().GetText()
	if ctx.Identifier() == nil {
		return []string{w}
	}
	identifier := v.Visit(ctx.Identifier()).([]string)
	identifier = append(identifier, w)
	return identifier
}

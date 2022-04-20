// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by InformerParser.
type InformerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by InformerParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by InformerParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by InformerParser#Rulebook.
	VisitRulebook(ctx *RulebookContext) interface{}

	// Visit a parse tree produced by InformerParser#Activity.
	VisitActivity(ctx *ActivityContext) interface{}

	// Visit a parse tree produced by InformerParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}

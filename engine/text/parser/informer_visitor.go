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

	// Visit a parse tree produced by InformerParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by InformerParser#Rulebook.
	VisitRulebook(ctx *RulebookContext) interface{}

	// Visit a parse tree produced by InformerParser#Instanciate.
	VisitInstanciate(ctx *InstanciateContext) interface{}

	// Visit a parse tree produced by InformerParser#ObjectBasedRulebook.
	VisitObjectBasedRulebook(ctx *ObjectBasedRulebookContext) interface{}

	// Visit a parse tree produced by InformerParser#Activity.
	VisitActivity(ctx *ActivityContext) interface{}

	// Visit a parse tree produced by InformerParser#CertaintyOfAttribute.
	VisitCertaintyOfAttribute(ctx *CertaintyOfAttributeContext) interface{}

	// Visit a parse tree produced by InformerParser#ObjectKind.
	VisitObjectKind(ctx *ObjectKindContext) interface{}

	// Visit a parse tree produced by InformerParser#ValueKind.
	VisitValueKind(ctx *ValueKindContext) interface{}

	// Visit a parse tree produced by InformerParser#certainty.
	VisitCertainty(ctx *CertaintyContext) interface{}

	// Visit a parse tree produced by InformerParser#designator.
	VisitDesignator(ctx *DesignatorContext) interface{}

	// Visit a parse tree produced by InformerParser#values.
	VisitValues(ctx *ValuesContext) interface{}
}

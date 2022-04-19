// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInformerListener is a complete listener for a parse tree produced by InformerParser.
type BaseInformerListener struct{}

var _ InformerListener = &BaseInformerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInformerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInformerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInformerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInformerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseInformerListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseInformerListener) ExitStart(ctx *StartContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseInformerListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInformerListener) ExitStatement(ctx *StatementContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseInformerListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseInformerListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterRulebook is called when production Rulebook is entered.
func (s *BaseInformerListener) EnterRulebook(ctx *RulebookContext) {}

// ExitRulebook is called when production Rulebook is exited.
func (s *BaseInformerListener) ExitRulebook(ctx *RulebookContext) {}

// EnterObjectBasedRulebook is called when production ObjectBasedRulebook is entered.
func (s *BaseInformerListener) EnterObjectBasedRulebook(ctx *ObjectBasedRulebookContext) {}

// ExitObjectBasedRulebook is called when production ObjectBasedRulebook is exited.
func (s *BaseInformerListener) ExitObjectBasedRulebook(ctx *ObjectBasedRulebookContext) {}

// EnterActivity is called when production Activity is entered.
func (s *BaseInformerListener) EnterActivity(ctx *ActivityContext) {}

// ExitActivity is called when production Activity is exited.
func (s *BaseInformerListener) ExitActivity(ctx *ActivityContext) {}

// EnterCertaintyOfAttribute is called when production CertaintyOfAttribute is entered.
func (s *BaseInformerListener) EnterCertaintyOfAttribute(ctx *CertaintyOfAttributeContext) {}

// ExitCertaintyOfAttribute is called when production CertaintyOfAttribute is exited.
func (s *BaseInformerListener) ExitCertaintyOfAttribute(ctx *CertaintyOfAttributeContext) {}

// EnterObjectKind is called when production ObjectKind is entered.
func (s *BaseInformerListener) EnterObjectKind(ctx *ObjectKindContext) {}

// ExitObjectKind is called when production ObjectKind is exited.
func (s *BaseInformerListener) ExitObjectKind(ctx *ObjectKindContext) {}

// EnterValueKind is called when production ValueKind is entered.
func (s *BaseInformerListener) EnterValueKind(ctx *ValueKindContext) {}

// ExitValueKind is called when production ValueKind is exited.
func (s *BaseInformerListener) ExitValueKind(ctx *ValueKindContext) {}

// EnterCertainty is called when production certainty is entered.
func (s *BaseInformerListener) EnterCertainty(ctx *CertaintyContext) {}

// ExitCertainty is called when production certainty is exited.
func (s *BaseInformerListener) ExitCertainty(ctx *CertaintyContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseInformerListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseInformerListener) ExitDesignator(ctx *DesignatorContext) {}

// EnterValues is called when production values is entered.
func (s *BaseInformerListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseInformerListener) ExitValues(ctx *ValuesContext) {}

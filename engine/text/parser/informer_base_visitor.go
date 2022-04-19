// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseInformerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseInformerVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitRulebook(ctx *RulebookContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitInstanciate(ctx *InstanciateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitObjectBasedRulebook(ctx *ObjectBasedRulebookContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitActivity(ctx *ActivityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitCertaintyOfAttribute(ctx *CertaintyOfAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitObjectKind(ctx *ObjectKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitValueKind(ctx *ValueKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitCertainty(ctx *CertaintyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitDesignator(ctx *DesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseInformerVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

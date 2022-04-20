// Code generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Informer
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type InformerParser struct {
	*antlr.BaseParser
}

var informerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func informerParserInit() {
	staticData := &informerParserStaticData
	staticData.literalNames = []string{
		"", "'.'", "'is a rulebook'", "'is an object based rulebook'", "'is an activity'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "WORD", "PUNCT", "COMMENT", "WHITESPACE", "EOL",
	}
	staticData.ruleNames = []string{
		"start", "statement", "definition", "identifier",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 36, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 4, 0,
		10, 8, 0, 11, 0, 12, 0, 11, 1, 1, 1, 1, 3, 1, 16, 8, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 29, 8, 2, 1, 3,
		1, 3, 1, 3, 3, 3, 34, 8, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 36, 0, 9,
		1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8,
		10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0,
		11, 12, 1, 0, 0, 0, 12, 1, 1, 0, 0, 0, 13, 15, 3, 4, 2, 0, 14, 16, 5, 1,
		0, 0, 15, 14, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18,
		5, 9, 0, 0, 18, 3, 1, 0, 0, 0, 19, 20, 3, 6, 3, 0, 20, 21, 5, 2, 0, 0,
		21, 29, 1, 0, 0, 0, 22, 23, 3, 6, 3, 0, 23, 24, 5, 3, 0, 0, 24, 29, 1,
		0, 0, 0, 25, 26, 3, 6, 3, 0, 26, 27, 5, 4, 0, 0, 27, 29, 1, 0, 0, 0, 28,
		19, 1, 0, 0, 0, 28, 22, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29, 5, 1, 0, 0,
		0, 30, 34, 5, 5, 0, 0, 31, 32, 5, 5, 0, 0, 32, 34, 3, 6, 3, 0, 33, 30,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 7, 1, 0, 0, 0, 4, 11, 15, 28, 33,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// InformerParserInit initializes any static state used to implement InformerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInformerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InformerParserInit() {
	staticData := &informerParserStaticData
	staticData.once.Do(informerParserInit)
}

// NewInformerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInformerParser(input antlr.TokenStream) *InformerParser {
	InformerParserInit()
	this := new(InformerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &informerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Informer.g4"

	return this
}

// InformerParser tokens.
const (
	InformerParserEOF        = antlr.TokenEOF
	InformerParserT__0       = 1
	InformerParserT__1       = 2
	InformerParserT__2       = 3
	InformerParserT__3       = 4
	InformerParserWORD       = 5
	InformerParserPUNCT      = 6
	InformerParserCOMMENT    = 7
	InformerParserWHITESPACE = 8
	InformerParserEOL        = 9
)

// InformerParser rules.
const (
	InformerParserRULE_start      = 0
	InformerParserRULE_statement  = 1
	InformerParserRULE_definition = 2
	InformerParserRULE_identifier = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InformerVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InformerParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InformerParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == InformerParserWORD {
		{
			p.SetState(8)
			p.Statement()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *StatementContext) EOL() antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InformerVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InformerParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InformerParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Definition()
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InformerParserT__0 {
		{
			p.SetState(14)
			p.Match(InformerParserT__0)
		}

	}
	{
		p.SetState(17)
		p.Match(InformerParserEOL)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) CopyFrom(ctx *DefinitionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RulebookContext struct {
	*DefinitionContext
}

func NewRulebookContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RulebookContext {
	var p = new(RulebookContext)

	p.DefinitionContext = NewEmptyDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionContext))

	return p
}

func (s *RulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RulebookContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InformerVisitor:
		return t.VisitRulebook(s)

	default:
		return t.VisitChildren(s)
	}
}

type ActivityContext struct {
	*DefinitionContext
}

func NewActivityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ActivityContext {
	var p = new(ActivityContext)

	p.DefinitionContext = NewEmptyDefinitionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DefinitionContext))

	return p
}

func (s *ActivityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivityContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ActivityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InformerVisitor:
		return t.VisitActivity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InformerParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InformerParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Identifier()
		}
		{
			p.SetState(20)
			p.Match(InformerParserT__1)
		}

	case 2:
		localctx = NewRulebookContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Identifier()
		}
		{
			p.SetState(23)
			p.Match(InformerParserT__2)
		}

	case 3:
		localctx = NewActivityContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.Identifier()
		}
		{
			p.SetState(26)
			p.Match(InformerParserT__3)
		}

	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *IdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InformerVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InformerParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(InformerParserWORD)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(InformerParserWORD)
		}
		{
			p.SetState(32)
			p.Identifier()
		}

	}

	return localctx
}

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
		"", "'.'", "'is'", "'rulebook'", "'based'", "'activity'", "'kind'",
		"'of'", "'usually'", "'always'", "'never'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "ARTICLE", "GERUND",
		"WORD", "PUNCT", "COMMENT", "WHITESPACE", "EOL",
	}
	staticData.ruleNames = []string{
		"start", "statement", "definition", "definitionType", "certainty", "designator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 68, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 3, 0, 15, 8, 0, 1, 0, 1, 0, 4, 0, 19, 8, 0,
		11, 0, 12, 0, 20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 56, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 63, 8, 5, 10, 5,
		12, 5, 66, 9, 5, 1, 5, 0, 1, 10, 6, 0, 2, 4, 6, 8, 10, 0, 1, 1, 0, 8, 10,
		72, 0, 18, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 26, 1, 0, 0, 0, 6, 46, 1,
		0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 55, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13,
		15, 5, 1, 0, 0, 14, 13, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 16, 1, 0, 0,
		0, 16, 17, 5, 18, 0, 0, 17, 19, 1, 0, 0, 0, 18, 12, 1, 0, 0, 0, 19, 20,
		1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0,
		22, 23, 5, 0, 0, 1, 23, 1, 1, 0, 0, 0, 24, 25, 3, 4, 2, 0, 25, 3, 1, 0,
		0, 0, 26, 27, 3, 10, 5, 0, 27, 28, 5, 2, 0, 0, 28, 29, 3, 6, 3, 0, 29,
		5, 1, 0, 0, 0, 30, 31, 5, 12, 0, 0, 31, 47, 5, 3, 0, 0, 32, 33, 5, 12,
		0, 0, 33, 34, 3, 10, 5, 0, 34, 35, 5, 4, 0, 0, 35, 36, 5, 3, 0, 0, 36,
		47, 1, 0, 0, 0, 37, 38, 5, 12, 0, 0, 38, 47, 5, 5, 0, 0, 39, 40, 3, 8,
		4, 0, 40, 41, 3, 10, 5, 0, 41, 47, 1, 0, 0, 0, 42, 43, 5, 12, 0, 0, 43,
		44, 5, 6, 0, 0, 44, 45, 5, 7, 0, 0, 45, 47, 3, 10, 5, 0, 46, 30, 1, 0,
		0, 0, 46, 32, 1, 0, 0, 0, 46, 37, 1, 0, 0, 0, 46, 39, 1, 0, 0, 0, 46, 42,
		1, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 49, 7, 0, 0, 0, 49, 9, 1, 0, 0, 0, 50,
		51, 6, 5, -1, 0, 51, 56, 5, 14, 0, 0, 52, 56, 5, 12, 0, 0, 53, 56, 5, 13,
		0, 0, 54, 56, 5, 7, 0, 0, 55, 50, 1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 55, 53,
		1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 64, 1, 0, 0, 0, 57, 58, 10, 2, 0, 0,
		58, 59, 5, 11, 0, 0, 59, 63, 3, 10, 5, 3, 60, 61, 10, 1, 0, 0, 61, 63,
		3, 10, 5, 2, 62, 57, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0,
		64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 11, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 6, 14, 20, 46, 55, 62, 64,
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
	InformerParserT__4       = 5
	InformerParserT__5       = 6
	InformerParserT__6       = 7
	InformerParserT__7       = 8
	InformerParserT__8       = 9
	InformerParserT__9       = 10
	InformerParserT__10      = 11
	InformerParserARTICLE    = 12
	InformerParserGERUND     = 13
	InformerParserWORD       = 14
	InformerParserPUNCT      = 15
	InformerParserCOMMENT    = 16
	InformerParserWHITESPACE = 17
	InformerParserEOL        = 18
)

// InformerParser rules.
const (
	InformerParserRULE_start          = 0
	InformerParserRULE_statement      = 1
	InformerParserRULE_definition     = 2
	InformerParserRULE_definitionType = 3
	InformerParserRULE_certainty      = 4
	InformerParserRULE_designator     = 5
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

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(InformerParserEOF, 0)
}

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

func (s *StartContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(InformerParserEOL)
}

func (s *StartContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(InformerParserEOL, i)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitStart(s)
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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InformerParserT__6)|(1<<InformerParserARTICLE)|(1<<InformerParserGERUND)|(1<<InformerParserWORD))) != 0) {
		{
			p.SetState(12)
			p.Statement()
		}
		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == InformerParserT__0 {
			{
				p.SetState(13)
				p.Match(InformerParserT__0)
			}

		}
		{
			p.SetState(16)
			p.Match(InformerParserEOL)
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(InformerParserEOF)
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

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *InformerParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InformerParserRULE_statement)

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
		p.SetState(24)
		p.Definition()
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

func (s *DefinitionContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *DefinitionContext) DefinitionType() IDefinitionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionTypeContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitDefinition(s)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.designator(0)
	}
	{
		p.SetState(27)
		p.Match(InformerParserT__1)
	}
	{
		p.SetState(28)
		p.DefinitionType()
	}

	return localctx
}

// IDefinitionTypeContext is an interface to support dynamic dispatch.
type IDefinitionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionTypeContext differentiates from other interfaces.
	IsDefinitionTypeContext()
}

type DefinitionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionTypeContext() *DefinitionTypeContext {
	var p = new(DefinitionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_definitionType
	return p
}

func (*DefinitionTypeContext) IsDefinitionTypeContext() {}

func NewDefinitionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionTypeContext {
	var p = new(DefinitionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_definitionType

	return p
}

func (s *DefinitionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionTypeContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *DefinitionTypeContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *DefinitionTypeContext) Certainty() ICertaintyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICertaintyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICertaintyContext)
}

func (s *DefinitionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterDefinitionType(s)
	}
}

func (s *DefinitionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitDefinitionType(s)
	}
}

func (p *InformerParser) DefinitionType() (localctx IDefinitionTypeContext) {
	this := p
	_ = this

	localctx = NewDefinitionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_definitionType)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(31)
			p.Match(InformerParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(33)
			p.designator(0)
		}
		{
			p.SetState(34)
			p.Match(InformerParserT__3)
		}
		{
			p.SetState(35)
			p.Match(InformerParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(38)
			p.Match(InformerParserT__4)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Certainty()
		}
		{
			p.SetState(40)
			p.designator(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(43)
			p.Match(InformerParserT__5)
		}
		{
			p.SetState(44)
			p.Match(InformerParserT__6)
		}
		{
			p.SetState(45)
			p.designator(0)
		}

	}

	return localctx
}

// ICertaintyContext is an interface to support dynamic dispatch.
type ICertaintyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCertaintyContext differentiates from other interfaces.
	IsCertaintyContext()
}

type CertaintyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCertaintyContext() *CertaintyContext {
	var p = new(CertaintyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_certainty
	return p
}

func (*CertaintyContext) IsCertaintyContext() {}

func NewCertaintyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CertaintyContext {
	var p = new(CertaintyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_certainty

	return p
}

func (s *CertaintyContext) GetParser() antlr.Parser { return s.parser }
func (s *CertaintyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CertaintyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CertaintyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterCertainty(s)
	}
}

func (s *CertaintyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitCertainty(s)
	}
}

func (p *InformerParser) Certainty() (localctx ICertaintyContext) {
	this := p
	_ = this

	localctx = NewCertaintyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InformerParserRULE_certainty)
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
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InformerParserT__7)|(1<<InformerParserT__8)|(1<<InformerParserT__9))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDesignatorContext is an interface to support dynamic dispatch.
type IDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesignatorContext differentiates from other interfaces.
	IsDesignatorContext()
}

type DesignatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorContext() *DesignatorContext {
	var p = new(DesignatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_designator
	return p
}

func (*DesignatorContext) IsDesignatorContext() {}

func NewDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorContext {
	var p = new(DesignatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_designator

	return p
}

func (s *DesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorContext) WORD() antlr.TerminalNode {
	return s.GetToken(InformerParserWORD, 0)
}

func (s *DesignatorContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *DesignatorContext) GERUND() antlr.TerminalNode {
	return s.GetToken(InformerParserGERUND, 0)
}

func (s *DesignatorContext) AllDesignator() []IDesignatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDesignatorContext); ok {
			len++
		}
	}

	tst := make([]IDesignatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDesignatorContext); ok {
			tst[i] = t.(IDesignatorContext)
			i++
		}
	}

	return tst
}

func (s *DesignatorContext) Designator(i int) IDesignatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
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

	return t.(IDesignatorContext)
}

func (s *DesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterDesignator(s)
	}
}

func (s *DesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitDesignator(s)
	}
}

func (p *InformerParser) Designator() (localctx IDesignatorContext) {
	return p.designator(0)
}

func (p *InformerParser) designator(_p int) (localctx IDesignatorContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDesignatorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, InformerParserRULE_designator, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InformerParserWORD:
		{
			p.SetState(51)
			p.Match(InformerParserWORD)
		}

	case InformerParserARTICLE:
		{
			p.SetState(52)
			p.Match(InformerParserARTICLE)
		}

	case InformerParserGERUND:
		{
			p.SetState(53)
			p.Match(InformerParserGERUND)
		}

	case InformerParserT__6:
		{
			p.SetState(54)
			p.Match(InformerParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(58)
					p.Match(InformerParserT__10)
				}
				{
					p.SetState(59)
					p.designator(3)
				}

			case 2:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(61)
					p.designator(2)
				}

			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

func (p *InformerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *DesignatorContext = nil
		if localctx != nil {
			t = localctx.(*DesignatorContext)
		}
		return p.Designator_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *InformerParser) Designator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

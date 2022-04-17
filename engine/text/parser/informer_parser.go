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
		"", "'.'", "'is'", "'rulebook'", "'based'", "'activity'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "ARTICLE", "WORD", "PUNCT", "COMMENT", "WHITESPACE",
		"EOL",
	}
	staticData.ruleNames = []string{
		"start", "statement", "rulebook", "activity", "designator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 61, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 3, 0, 13, 8, 0, 1, 0, 1, 0, 4, 0, 17, 8, 0, 11, 0, 12, 0,
		18, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 25, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 49, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1, 4, 0, 1, 8, 5, 0, 2,
		4, 6, 8, 0, 0, 62, 0, 16, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 38, 1, 0, 0,
		0, 6, 40, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 13, 5,
		1, 0, 0, 12, 11, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14,
		15, 5, 12, 0, 0, 15, 17, 1, 0, 0, 0, 16, 10, 1, 0, 0, 0, 17, 18, 1, 0,
		0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21,
		5, 0, 0, 1, 21, 1, 1, 0, 0, 0, 22, 25, 3, 4, 2, 0, 23, 25, 3, 6, 3, 0,
		24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 3, 1, 0, 0, 0, 26, 27, 3, 8,
		4, 0, 27, 28, 5, 2, 0, 0, 28, 29, 5, 7, 0, 0, 29, 30, 5, 3, 0, 0, 30, 39,
		1, 0, 0, 0, 31, 32, 3, 8, 4, 0, 32, 33, 5, 2, 0, 0, 33, 34, 5, 7, 0, 0,
		34, 35, 3, 8, 4, 0, 35, 36, 5, 4, 0, 0, 36, 37, 5, 3, 0, 0, 37, 39, 1,
		0, 0, 0, 38, 26, 1, 0, 0, 0, 38, 31, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40,
		41, 3, 8, 4, 0, 41, 42, 5, 2, 0, 0, 42, 43, 5, 7, 0, 0, 43, 44, 5, 5, 0,
		0, 44, 7, 1, 0, 0, 0, 45, 46, 6, 4, -1, 0, 46, 49, 5, 8, 0, 0, 47, 49,
		5, 7, 0, 0, 48, 45, 1, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 57, 1, 0, 0, 0,
		50, 51, 10, 2, 0, 0, 51, 52, 5, 6, 0, 0, 52, 56, 3, 8, 4, 3, 53, 54, 10,
		1, 0, 0, 54, 56, 3, 8, 4, 2, 55, 50, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56,
		59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 9, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 7, 12, 18, 24, 38, 48, 55, 57,
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
	InformerParserARTICLE    = 7
	InformerParserWORD       = 8
	InformerParserPUNCT      = 9
	InformerParserCOMMENT    = 10
	InformerParserWHITESPACE = 11
	InformerParserEOL        = 12
)

// InformerParser rules.
const (
	InformerParserRULE_start      = 0
	InformerParserRULE_statement  = 1
	InformerParserRULE_rulebook   = 2
	InformerParserRULE_activity   = 3
	InformerParserRULE_designator = 4
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
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == InformerParserARTICLE || _la == InformerParserWORD {
		{
			p.SetState(10)
			p.Statement()
		}
		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == InformerParserT__0 {
			{
				p.SetState(11)
				p.Match(InformerParserT__0)
			}

		}
		{
			p.SetState(14)
			p.Match(InformerParserEOL)
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
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

func (s *StatementContext) Rulebook() IRulebookContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulebookContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulebookContext)
}

func (s *StatementContext) Activity() IActivityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActivityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActivityContext)
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

	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Rulebook()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Activity()
		}

	}

	return localctx
}

// IRulebookContext is an interface to support dynamic dispatch.
type IRulebookContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulebookContext differentiates from other interfaces.
	IsRulebookContext()
}

type RulebookContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulebookContext() *RulebookContext {
	var p = new(RulebookContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_rulebook
	return p
}

func (*RulebookContext) IsRulebookContext() {}

func NewRulebookContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulebookContext {
	var p = new(RulebookContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_rulebook

	return p
}

func (s *RulebookContext) GetParser() antlr.Parser { return s.parser }

func (s *RulebookContext) AllDesignator() []IDesignatorContext {
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

func (s *RulebookContext) Designator(i int) IDesignatorContext {
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

func (s *RulebookContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *RulebookContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulebookContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulebookContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterRulebook(s)
	}
}

func (s *RulebookContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitRulebook(s)
	}
}

func (p *InformerParser) Rulebook() (localctx IRulebookContext) {
	this := p
	_ = this

	localctx = NewRulebookContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InformerParserRULE_rulebook)

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

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
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
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(29)
			p.Match(InformerParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.designator(0)
		}
		{
			p.SetState(32)
			p.Match(InformerParserT__1)
		}
		{
			p.SetState(33)
			p.Match(InformerParserARTICLE)
		}
		{
			p.SetState(34)
			p.designator(0)
		}
		{
			p.SetState(35)
			p.Match(InformerParserT__3)
		}
		{
			p.SetState(36)
			p.Match(InformerParserT__2)
		}

	}

	return localctx
}

// IActivityContext is an interface to support dynamic dispatch.
type IActivityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActivityContext differentiates from other interfaces.
	IsActivityContext()
}

type ActivityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActivityContext() *ActivityContext {
	var p = new(ActivityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InformerParserRULE_activity
	return p
}

func (*ActivityContext) IsActivityContext() {}

func NewActivityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActivityContext {
	var p = new(ActivityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformerParserRULE_activity

	return p
}

func (s *ActivityContext) GetParser() antlr.Parser { return s.parser }

func (s *ActivityContext) Designator() IDesignatorContext {
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

func (s *ActivityContext) ARTICLE() antlr.TerminalNode {
	return s.GetToken(InformerParserARTICLE, 0)
}

func (s *ActivityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActivityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.EnterActivity(s)
	}
}

func (s *ActivityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformerListener); ok {
		listenerT.ExitActivity(s)
	}
}

func (p *InformerParser) Activity() (localctx IActivityContext) {
	this := p
	_ = this

	localctx = NewActivityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformerParserRULE_activity)

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
		p.SetState(40)
		p.designator(0)
	}
	{
		p.SetState(41)
		p.Match(InformerParserT__1)
	}
	{
		p.SetState(42)
		p.Match(InformerParserARTICLE)
	}
	{
		p.SetState(43)
		p.Match(InformerParserT__4)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, InformerParserRULE_designator, _p)

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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InformerParserWORD:
		{
			p.SetState(46)
			p.Match(InformerParserWORD)
		}

	case InformerParserARTICLE:
		{
			p.SetState(47)
			p.Match(InformerParserARTICLE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(51)
					p.Match(InformerParserT__5)
				}
				{
					p.SetState(52)
					p.designator(3)
				}

			case 2:
				localctx = NewDesignatorContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, InformerParserRULE_designator)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(54)
					p.designator(2)
				}

			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *InformerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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

// Generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class InformerParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, NUMBER=14, PUNCT=15, STRING=16, 
		COMMENT=17, WHITESPACE=18, EOL=19, WORD=20, ANY=21;
	public static final int
		RULE_start = 0, RULE_definition = 1, RULE_title = 2, RULE_rulebookDeclaration = 3, 
		RULE_activityDeclaration = 4, RULE_actionDeclaration = 5, RULE_aliasDeclaration = 6, 
		RULE_designator = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "definition", "title", "rulebookDeclaration", "activityDeclaration", 
			"actionDeclaration", "aliasDeclaration", "designator"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'is'", "'a'", "'rulebook'", "'.'", "'an'", "'based'", "'activity'", 
			"'action'", "'applying'", "'to'", "'Understand'", "'as'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, "NUMBER", "PUNCT", "STRING", "COMMENT", "WHITESPACE", "EOL", 
			"WORD", "ANY"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Informer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public InformerParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(InformerParser.EOF, 0); }
		public List<DefinitionContext> definition() {
			return getRuleContexts(DefinitionContext.class);
		}
		public DefinitionContext definition(int i) {
			return getRuleContext(DefinitionContext.class,i);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(17); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(16);
				definition();
				}
				}
				setState(19); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__10) | (1L << STRING) | (1L << WORD))) != 0) );
			setState(21);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DefinitionContext extends ParserRuleContext {
		public ActivityDeclarationContext activityDeclaration() {
			return getRuleContext(ActivityDeclarationContext.class,0);
		}
		public RulebookDeclarationContext rulebookDeclaration() {
			return getRuleContext(RulebookDeclarationContext.class,0);
		}
		public ActionDeclarationContext actionDeclaration() {
			return getRuleContext(ActionDeclarationContext.class,0);
		}
		public AliasDeclarationContext aliasDeclaration() {
			return getRuleContext(AliasDeclarationContext.class,0);
		}
		public TitleContext title() {
			return getRuleContext(TitleContext.class,0);
		}
		public DefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definition; }
	}

	public final DefinitionContext definition() throws RecognitionException {
		DefinitionContext _localctx = new DefinitionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_definition);
		try {
			setState(28);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(23);
				activityDeclaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(24);
				rulebookDeclaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(25);
				actionDeclaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(26);
				aliasDeclaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(27);
				title();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TitleContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(InformerParser.STRING, 0); }
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public TitleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_title; }
	}

	public final TitleContext title() throws RecognitionException {
		TitleContext _localctx = new TitleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_title);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(STRING);
			setState(31);
			match(EOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RulebookDeclarationContext extends ParserRuleContext {
		public RulebookDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rulebookDeclaration; }
	 
		public RulebookDeclarationContext() { }
		public void copyFrom(RulebookDeclarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class GenericRulebookContext extends RulebookDeclarationContext {
		public DesignatorContext designator() {
			return getRuleContext(DesignatorContext.class,0);
		}
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public GenericRulebookContext(RulebookDeclarationContext ctx) { copyFrom(ctx); }
	}
	public static class SpecificRulebookContext extends RulebookDeclarationContext {
		public DesignatorContext designator() {
			return getRuleContext(DesignatorContext.class,0);
		}
		public TerminalNode WORD() { return getToken(InformerParser.WORD, 0); }
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public SpecificRulebookContext(RulebookDeclarationContext ctx) { copyFrom(ctx); }
	}

	public final RulebookDeclarationContext rulebookDeclaration() throws RecognitionException {
		RulebookDeclarationContext _localctx = new RulebookDeclarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_rulebookDeclaration);
		int _la;
		try {
			setState(49);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new GenericRulebookContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				designator(0);
				setState(34);
				match(T__0);
				setState(35);
				match(T__1);
				setState(36);
				match(T__2);
				setState(37);
				match(T__3);
				setState(38);
				match(EOL);
				}
				break;
			case 2:
				_localctx = new SpecificRulebookContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(40);
				designator(0);
				setState(41);
				match(T__0);
				setState(42);
				_la = _input.LA(1);
				if ( !(_la==T__1 || _la==T__4) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(43);
				match(WORD);
				setState(44);
				match(T__5);
				setState(45);
				match(T__2);
				setState(46);
				match(T__3);
				setState(47);
				match(EOL);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ActivityDeclarationContext extends ParserRuleContext {
		public DesignatorContext designator() {
			return getRuleContext(DesignatorContext.class,0);
		}
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public ActivityDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_activityDeclaration; }
	}

	public final ActivityDeclarationContext activityDeclaration() throws RecognitionException {
		ActivityDeclarationContext _localctx = new ActivityDeclarationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_activityDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			designator(0);
			setState(52);
			match(T__0);
			setState(53);
			match(T__4);
			setState(54);
			match(T__6);
			setState(55);
			match(T__3);
			setState(56);
			match(EOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ActionDeclarationContext extends ParserRuleContext {
		public DesignatorContext designator() {
			return getRuleContext(DesignatorContext.class,0);
		}
		public TerminalNode WORD() { return getToken(InformerParser.WORD, 0); }
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public ActionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_actionDeclaration; }
	}

	public final ActionDeclarationContext actionDeclaration() throws RecognitionException {
		ActionDeclarationContext _localctx = new ActionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_actionDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			designator(0);
			setState(59);
			match(T__0);
			setState(60);
			match(T__4);
			setState(61);
			match(T__7);
			setState(62);
			match(T__8);
			setState(63);
			match(T__9);
			setState(64);
			match(T__1);
			setState(65);
			match(WORD);
			setState(66);
			match(T__3);
			setState(67);
			match(EOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AliasDeclarationContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(InformerParser.STRING, 0); }
		public TerminalNode WORD() { return getToken(InformerParser.WORD, 0); }
		public TerminalNode EOL() { return getToken(InformerParser.EOL, 0); }
		public AliasDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aliasDeclaration; }
	}

	public final AliasDeclarationContext aliasDeclaration() throws RecognitionException {
		AliasDeclarationContext _localctx = new AliasDeclarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_aliasDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(T__10);
			setState(70);
			match(STRING);
			setState(71);
			match(T__11);
			setState(72);
			match(WORD);
			setState(73);
			match(T__3);
			setState(74);
			match(EOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DesignatorContext extends ParserRuleContext {
		public TerminalNode WORD() { return getToken(InformerParser.WORD, 0); }
		public DesignatorContext designator() {
			return getRuleContext(DesignatorContext.class,0);
		}
		public DesignatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_designator; }
	}

	public final DesignatorContext designator() throws RecognitionException {
		return designator(0);
	}

	private DesignatorContext designator(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DesignatorContext _localctx = new DesignatorContext(_ctx, _parentState);
		DesignatorContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_designator, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(77);
			match(WORD);
			}
			_ctx.stop = _input.LT(-1);
			setState(86);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(84);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new DesignatorContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_designator);
						setState(79);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(80);
						match(WORD);
						}
						break;
					case 2:
						{
						_localctx = new DesignatorContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_designator);
						setState(81);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(82);
						match(T__12);
						setState(83);
						match(WORD);
						}
						break;
					}
					} 
				}
				setState(88);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return designator_sempred((DesignatorContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean designator_sempred(DesignatorContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\27\\\4\2\t\2\4\3"+
		"\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\6\2\24\n\2\r"+
		"\2\16\2\25\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5\3\37\n\3\3\4\3\4\3\4\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\64\n\5\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\tW\n\t"+
		"\f\t\16\tZ\13\t\3\t\2\3\20\n\2\4\6\b\n\f\16\20\2\3\4\2\4\4\7\7\2[\2\23"+
		"\3\2\2\2\4\36\3\2\2\2\6 \3\2\2\2\b\63\3\2\2\2\n\65\3\2\2\2\f<\3\2\2\2"+
		"\16G\3\2\2\2\20N\3\2\2\2\22\24\5\4\3\2\23\22\3\2\2\2\24\25\3\2\2\2\25"+
		"\23\3\2\2\2\25\26\3\2\2\2\26\27\3\2\2\2\27\30\7\2\2\3\30\3\3\2\2\2\31"+
		"\37\5\n\6\2\32\37\5\b\5\2\33\37\5\f\7\2\34\37\5\16\b\2\35\37\5\6\4\2\36"+
		"\31\3\2\2\2\36\32\3\2\2\2\36\33\3\2\2\2\36\34\3\2\2\2\36\35\3\2\2\2\37"+
		"\5\3\2\2\2 !\7\22\2\2!\"\7\25\2\2\"\7\3\2\2\2#$\5\20\t\2$%\7\3\2\2%&\7"+
		"\4\2\2&\'\7\5\2\2\'(\7\6\2\2()\7\25\2\2)\64\3\2\2\2*+\5\20\t\2+,\7\3\2"+
		"\2,-\t\2\2\2-.\7\26\2\2./\7\b\2\2/\60\7\5\2\2\60\61\7\6\2\2\61\62\7\25"+
		"\2\2\62\64\3\2\2\2\63#\3\2\2\2\63*\3\2\2\2\64\t\3\2\2\2\65\66\5\20\t\2"+
		"\66\67\7\3\2\2\678\7\7\2\289\7\t\2\29:\7\6\2\2:;\7\25\2\2;\13\3\2\2\2"+
		"<=\5\20\t\2=>\7\3\2\2>?\7\7\2\2?@\7\n\2\2@A\7\13\2\2AB\7\f\2\2BC\7\4\2"+
		"\2CD\7\26\2\2DE\7\6\2\2EF\7\25\2\2F\r\3\2\2\2GH\7\r\2\2HI\7\22\2\2IJ\7"+
		"\16\2\2JK\7\26\2\2KL\7\6\2\2LM\7\25\2\2M\17\3\2\2\2NO\b\t\1\2OP\7\26\2"+
		"\2PX\3\2\2\2QR\f\5\2\2RW\7\26\2\2ST\f\4\2\2TU\7\17\2\2UW\7\26\2\2VQ\3"+
		"\2\2\2VS\3\2\2\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\21\3\2\2\2ZX\3\2\2\2\7"+
		"\25\36\63VX";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
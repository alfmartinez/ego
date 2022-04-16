// Generated from /Users/alex/Projects/ego/engine/text/Informer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class InformerLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, NUMBER=5, IDENT=6, PUNCT=7, STRING=8, 
		COMMENT=9, WHITESPACE=10;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "NUMBER", "IDENT", "PUNCT", "STRING", 
			"COMMENT", "WHITESPACE", "ESC"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'is'", "'a'", "'rulebook'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "NUMBER", "IDENT", "PUNCT", "STRING", "COMMENT", 
			"WHITESPACE"
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


	public InformerLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Informer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\fW\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\5\3\5\3\6\6\6+\n\6\r\6\16\6,\3\7\6\7\60\n\7\r\7\16\7\61\3\b\3\b\3\t\3"+
		"\t\3\t\7\t9\n\t\f\t\16\t<\13\t\3\t\3\t\3\n\3\n\7\nB\n\n\f\n\16\nE\13\n"+
		"\3\n\3\n\3\n\3\n\3\13\6\13L\n\13\r\13\16\13M\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\5\fV\n\f\4:C\2\r\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\2\3"+
		"\2\6\3\2\62;\4\2C\\c|\6\2..\60\60<<^^\5\2\13\f\17\17\"\"\2\\\2\3\3\2\2"+
		"\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3"+
		"\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\3\31\3\2\2\2\5\34\3\2\2"+
		"\2\7\36\3\2\2\2\t\'\3\2\2\2\13*\3\2\2\2\r/\3\2\2\2\17\63\3\2\2\2\21\65"+
		"\3\2\2\2\23?\3\2\2\2\25K\3\2\2\2\27U\3\2\2\2\31\32\7k\2\2\32\33\7u\2\2"+
		"\33\4\3\2\2\2\34\35\7c\2\2\35\6\3\2\2\2\36\37\7t\2\2\37 \7w\2\2 !\7n\2"+
		"\2!\"\7g\2\2\"#\7d\2\2#$\7q\2\2$%\7q\2\2%&\7m\2\2&\b\3\2\2\2\'(\7\60\2"+
		"\2(\n\3\2\2\2)+\t\2\2\2*)\3\2\2\2+,\3\2\2\2,*\3\2\2\2,-\3\2\2\2-\f\3\2"+
		"\2\2.\60\t\3\2\2/.\3\2\2\2\60\61\3\2\2\2\61/\3\2\2\2\61\62\3\2\2\2\62"+
		"\16\3\2\2\2\63\64\t\4\2\2\64\20\3\2\2\2\65:\7$\2\2\669\5\27\f\2\679\13"+
		"\2\2\28\66\3\2\2\28\67\3\2\2\29<\3\2\2\2:;\3\2\2\2:8\3\2\2\2;=\3\2\2\2"+
		"<:\3\2\2\2=>\7$\2\2>\22\3\2\2\2?C\7]\2\2@B\13\2\2\2A@\3\2\2\2BE\3\2\2"+
		"\2CD\3\2\2\2CA\3\2\2\2DF\3\2\2\2EC\3\2\2\2FG\7_\2\2GH\3\2\2\2HI\b\n\2"+
		"\2I\24\3\2\2\2JL\t\5\2\2KJ\3\2\2\2LM\3\2\2\2MK\3\2\2\2MN\3\2\2\2NO\3\2"+
		"\2\2OP\b\13\2\2P\26\3\2\2\2QR\7^\2\2RV\7$\2\2ST\7^\2\2TV\7^\2\2UQ\3\2"+
		"\2\2US\3\2\2\2V\30\3\2\2\2\n\2,\618:CMU\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
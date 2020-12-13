// Generated from /home/andreas/master/inf225/project/Garvik.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GarvikLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		BOOL=25, ID=26, NUM=27, STR=28, WS=29;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "BOOL", 
			"ID", "NUM", "STR", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'=='", "'<'", "'>'", "'.'", "'['", "']'", "'='", "'('", "')'", 
			"'/'", "'*'", "'-'", "'+'", "'->'", "'//'", "','", "'{'", "'}'", "'if'", 
			"'then'", "'else'", "'let'", "'in'", "'len'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "BOOL", "ID", "NUM", "STR", "WS"
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


	public GarvikLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Garvik.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\37\u00a3\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\3\2\3\2\3\2\3"+
		"\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3"+
		"\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21"+
		"\3\22\3\22\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u0086\n\32\3\33\6\33"+
		"\u0089\n\33\r\33\16\33\u008a\3\34\6\34\u008e\n\34\r\34\16\34\u008f\3\35"+
		"\3\35\3\35\3\35\7\35\u0096\n\35\f\35\16\35\u0099\13\35\3\35\3\35\3\36"+
		"\6\36\u009e\n\36\r\36\16\36\u009f\3\36\3\36\2\2\37\3\3\5\4\7\5\t\6\13"+
		"\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'"+
		"\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37\3\2\7\3\2c|\3\2\62"+
		";\n\2$$))^^ddhhppttvv\6\2\f\f\17\17$$^^\5\2\13\f\17\17\"\"\2\u00a8\2\3"+
		"\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2"+
		"\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31"+
		"\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2"+
		"\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2"+
		"\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2"+
		"\2\3=\3\2\2\2\5@\3\2\2\2\7B\3\2\2\2\tD\3\2\2\2\13F\3\2\2\2\rH\3\2\2\2"+
		"\17J\3\2\2\2\21L\3\2\2\2\23N\3\2\2\2\25P\3\2\2\2\27R\3\2\2\2\31T\3\2\2"+
		"\2\33V\3\2\2\2\35X\3\2\2\2\37[\3\2\2\2!^\3\2\2\2#`\3\2\2\2%b\3\2\2\2\'"+
		"d\3\2\2\2)g\3\2\2\2+l\3\2\2\2-q\3\2\2\2/u\3\2\2\2\61x\3\2\2\2\63\u0085"+
		"\3\2\2\2\65\u0088\3\2\2\2\67\u008d\3\2\2\29\u0091\3\2\2\2;\u009d\3\2\2"+
		"\2=>\7?\2\2>?\7?\2\2?\4\3\2\2\2@A\7>\2\2A\6\3\2\2\2BC\7@\2\2C\b\3\2\2"+
		"\2DE\7\60\2\2E\n\3\2\2\2FG\7]\2\2G\f\3\2\2\2HI\7_\2\2I\16\3\2\2\2JK\7"+
		"?\2\2K\20\3\2\2\2LM\7*\2\2M\22\3\2\2\2NO\7+\2\2O\24\3\2\2\2PQ\7\61\2\2"+
		"Q\26\3\2\2\2RS\7,\2\2S\30\3\2\2\2TU\7/\2\2U\32\3\2\2\2VW\7-\2\2W\34\3"+
		"\2\2\2XY\7/\2\2YZ\7@\2\2Z\36\3\2\2\2[\\\7\61\2\2\\]\7\61\2\2] \3\2\2\2"+
		"^_\7.\2\2_\"\3\2\2\2`a\7}\2\2a$\3\2\2\2bc\7\177\2\2c&\3\2\2\2de\7k\2\2"+
		"ef\7h\2\2f(\3\2\2\2gh\7v\2\2hi\7j\2\2ij\7g\2\2jk\7p\2\2k*\3\2\2\2lm\7"+
		"g\2\2mn\7n\2\2no\7u\2\2op\7g\2\2p,\3\2\2\2qr\7n\2\2rs\7g\2\2st\7v\2\2"+
		"t.\3\2\2\2uv\7k\2\2vw\7p\2\2w\60\3\2\2\2xy\7n\2\2yz\7g\2\2z{\7p\2\2{\62"+
		"\3\2\2\2|}\7v\2\2}~\7t\2\2~\177\7w\2\2\177\u0086\7g\2\2\u0080\u0081\7"+
		"h\2\2\u0081\u0082\7c\2\2\u0082\u0083\7n\2\2\u0083\u0084\7u\2\2\u0084\u0086"+
		"\7g\2\2\u0085|\3\2\2\2\u0085\u0080\3\2\2\2\u0086\64\3\2\2\2\u0087\u0089"+
		"\t\2\2\2\u0088\u0087\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u0088\3\2\2\2\u008a"+
		"\u008b\3\2\2\2\u008b\66\3\2\2\2\u008c\u008e\t\3\2\2\u008d\u008c\3\2\2"+
		"\2\u008e\u008f\3\2\2\2\u008f\u008d\3\2\2\2\u008f\u0090\3\2\2\2\u00908"+
		"\3\2\2\2\u0091\u0097\7$\2\2\u0092\u0093\7^\2\2\u0093\u0096\t\4\2\2\u0094"+
		"\u0096\n\5\2\2\u0095\u0092\3\2\2\2\u0095\u0094\3\2\2\2\u0096\u0099\3\2"+
		"\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u009a\3\2\2\2\u0099"+
		"\u0097\3\2\2\2\u009a\u009b\7$\2\2\u009b:\3\2\2\2\u009c\u009e\t\6\2\2\u009d"+
		"\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2"+
		"\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a2\b\36\2\2\u00a2<\3\2\2\2\t\2\u0085"+
		"\u008a\u008f\u0095\u0097\u009f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from /home/andreas/master/inf225/project/Garvik.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GarvikParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, BOOL=28, ID=29, NUM=30, STR=31, WS=32;
	public static final int
		RULE_program = 0, RULE_expr = 1;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'=='", "'<'", "'>'", "'.'", "'='", "'['", "']'", "'('", "')'", 
			"'/'", "'*'", "'-'", "'+'", "'->'", "':'", "'//'", "','", "'{'", "'}'", 
			"'if'", "'then'", "'else'", "'let'", "'in'", "'pop'", "'drop'", "'len'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "BOOL", "ID", "NUM", "STR", "WS"
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
	public String getGrammarFileName() { return "Garvik.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GarvikParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(7);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__5) | (1L << T__7) | (1L << T__11) | (1L << T__15) | (1L << T__17) | (1L << T__19) | (1L << T__22) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << BOOL) | (1L << ID) | (1L << NUM) | (1L << STR))) != 0)) {
				{
				{
				setState(4);
				expr(0);
				}
				}
				setState(9);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class StrExprContext extends ExprContext {
		public TerminalNode STR() { return getToken(GarvikParser.STR, 0); }
		public StrExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DotExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext field;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public DotExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LenExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public LenExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NegativeExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NegativeExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class PopExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PopExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class SubExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public SubExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ParenExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParenExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LetExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext value;
		public ExprContext expression;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LetExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class GreaterExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public GreaterExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LambdaExprContext extends ExprContext {
		public ExprContext param;
		public ExprContext body;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LambdaExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StructExprContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public StructExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DivExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public DivExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LookupAssignExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext key;
		public ExprContext value;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LookupAssignExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class CallExprContext extends ExprContext {
		public ExprContext fun;
		public ExprContext arg;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public CallExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ListExprContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ListExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AssignExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext value;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AssignExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IfElseExprContext extends ExprContext {
		public ExprContext con;
		public ExprContext t;
		public ExprContext f;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public IfElseExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LessExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LessExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DotAssignExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext field;
		public ExprContext value;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public DotAssignExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class MultExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MultExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class FieldExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext value;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public FieldExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class CommentExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public CommentExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NumExprContext extends ExprContext {
		public TerminalNode NUM() { return getToken(GarvikParser.NUM, 0); }
		public NumExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class DropExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DropExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AddExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AddExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class LookupExprContext extends ExprContext {
		public ExprContext id;
		public ExprContext key;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LookupExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class BoolExprContext extends ExprContext {
		public TerminalNode BOOL() { return getToken(GarvikParser.BOOL, 0); }
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GarvikParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class EqualExprContext extends ExprContext {
		public ExprContext left;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqualExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__11:
				{
				_localctx = new NegativeExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(11);
				match(T__11);
				setState(12);
				expr(14);
				}
				break;
			case T__15:
				{
				_localctx = new CommentExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(13);
				match(T__15);
				setState(14);
				expr(13);
				}
				break;
			case T__5:
				{
				_localctx = new ListExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(15);
				match(T__5);
				setState(16);
				expr(0);
				setState(21);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__16) {
					{
					{
					setState(17);
					match(T__16);
					setState(18);
					expr(0);
					}
					}
					setState(23);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(24);
				match(T__6);
				}
				break;
			case T__7:
				{
				_localctx = new ParenExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(26);
				match(T__7);
				setState(27);
				expr(0);
				setState(28);
				match(T__8);
				}
				break;
			case T__17:
				{
				_localctx = new StructExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(30);
				match(T__17);
				setState(34);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__5) | (1L << T__7) | (1L << T__11) | (1L << T__15) | (1L << T__17) | (1L << T__19) | (1L << T__22) | (1L << T__24) | (1L << T__25) | (1L << T__26) | (1L << BOOL) | (1L << ID) | (1L << NUM) | (1L << STR))) != 0)) {
					{
					{
					setState(31);
					expr(0);
					}
					}
					setState(36);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(37);
				match(T__18);
				}
				break;
			case T__19:
				{
				_localctx = new IfElseExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(38);
				match(T__19);
				setState(39);
				((IfElseExprContext)_localctx).con = expr(0);
				setState(40);
				match(T__20);
				setState(41);
				((IfElseExprContext)_localctx).t = expr(0);
				setState(42);
				match(T__21);
				setState(43);
				((IfElseExprContext)_localctx).f = expr(9);
				}
				break;
			case T__22:
				{
				_localctx = new LetExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(45);
				match(T__22);
				setState(46);
				((LetExprContext)_localctx).id = expr(0);
				setState(47);
				match(T__4);
				setState(48);
				((LetExprContext)_localctx).value = expr(0);
				setState(49);
				match(T__23);
				setState(50);
				((LetExprContext)_localctx).expression = expr(8);
				}
				break;
			case T__24:
				{
				_localctx = new PopExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(52);
				match(T__24);
				setState(53);
				((PopExprContext)_localctx).id = expr(7);
				}
				break;
			case T__25:
				{
				_localctx = new DropExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(54);
				match(T__25);
				setState(55);
				((DropExprContext)_localctx).id = expr(6);
				}
				break;
			case T__26:
				{
				_localctx = new LenExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(56);
				match(T__26);
				setState(57);
				((LenExprContext)_localctx).id = expr(5);
				}
				break;
			case BOOL:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(58);
				match(BOOL);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(59);
				match(ID);
				}
				break;
			case NUM:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(60);
				match(NUM);
				}
				break;
			case STR:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(61);
				match(STR);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(122);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(120);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new EqualExprContext(new ExprContext(_parentctx, _parentState));
						((EqualExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(64);
						if (!(precpred(_ctx, 29))) throw new FailedPredicateException(this, "precpred(_ctx, 29)");
						setState(65);
						match(T__0);
						setState(66);
						((EqualExprContext)_localctx).right = expr(30);
						}
						break;
					case 2:
						{
						_localctx = new LessExprContext(new ExprContext(_parentctx, _parentState));
						((LessExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(67);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(68);
						match(T__1);
						setState(69);
						((LessExprContext)_localctx).right = expr(29);
						}
						break;
					case 3:
						{
						_localctx = new GreaterExprContext(new ExprContext(_parentctx, _parentState));
						((GreaterExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(70);
						if (!(precpred(_ctx, 27))) throw new FailedPredicateException(this, "precpred(_ctx, 27)");
						setState(71);
						match(T__2);
						setState(72);
						((GreaterExprContext)_localctx).right = expr(28);
						}
						break;
					case 4:
						{
						_localctx = new DotAssignExprContext(new ExprContext(_parentctx, _parentState));
						((DotAssignExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(73);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(74);
						match(T__3);
						setState(75);
						((DotAssignExprContext)_localctx).field = expr(0);
						setState(76);
						match(T__4);
						setState(77);
						((DotAssignExprContext)_localctx).value = expr(27);
						}
						break;
					case 5:
						{
						_localctx = new DotExprContext(new ExprContext(_parentctx, _parentState));
						((DotExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(79);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(80);
						match(T__3);
						setState(81);
						((DotExprContext)_localctx).field = expr(26);
						}
						break;
					case 6:
						{
						_localctx = new LookupAssignExprContext(new ExprContext(_parentctx, _parentState));
						((LookupAssignExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(82);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(83);
						match(T__5);
						setState(84);
						((LookupAssignExprContext)_localctx).key = expr(0);
						setState(85);
						match(T__6);
						setState(86);
						match(T__4);
						setState(87);
						((LookupAssignExprContext)_localctx).value = expr(25);
						}
						break;
					case 7:
						{
						_localctx = new DivExprContext(new ExprContext(_parentctx, _parentState));
						((DivExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(89);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(90);
						match(T__9);
						setState(91);
						((DivExprContext)_localctx).right = expr(22);
						}
						break;
					case 8:
						{
						_localctx = new MultExprContext(new ExprContext(_parentctx, _parentState));
						((MultExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(92);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(93);
						match(T__10);
						setState(94);
						((MultExprContext)_localctx).right = expr(21);
						}
						break;
					case 9:
						{
						_localctx = new SubExprContext(new ExprContext(_parentctx, _parentState));
						((SubExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(95);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(96);
						match(T__11);
						setState(97);
						((SubExprContext)_localctx).right = expr(20);
						}
						break;
					case 10:
						{
						_localctx = new AddExprContext(new ExprContext(_parentctx, _parentState));
						((AddExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(98);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(99);
						match(T__12);
						setState(100);
						((AddExprContext)_localctx).right = expr(19);
						}
						break;
					case 11:
						{
						_localctx = new LambdaExprContext(new ExprContext(_parentctx, _parentState));
						((LambdaExprContext)_localctx).param = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(101);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(102);
						match(T__13);
						setState(103);
						((LambdaExprContext)_localctx).body = expr(18);
						}
						break;
					case 12:
						{
						_localctx = new FieldExprContext(new ExprContext(_parentctx, _parentState));
						((FieldExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(104);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(105);
						match(T__14);
						setState(106);
						((FieldExprContext)_localctx).value = expr(17);
						}
						break;
					case 13:
						{
						_localctx = new AssignExprContext(new ExprContext(_parentctx, _parentState));
						((AssignExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(107);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(108);
						match(T__4);
						setState(109);
						((AssignExprContext)_localctx).value = expr(16);
						}
						break;
					case 14:
						{
						_localctx = new LookupExprContext(new ExprContext(_parentctx, _parentState));
						((LookupExprContext)_localctx).id = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(110);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(111);
						match(T__5);
						setState(112);
						((LookupExprContext)_localctx).key = expr(0);
						setState(113);
						match(T__6);
						}
						break;
					case 15:
						{
						_localctx = new CallExprContext(new ExprContext(_parentctx, _parentState));
						((CallExprContext)_localctx).fun = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(115);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(116);
						match(T__7);
						setState(117);
						((CallExprContext)_localctx).arg = expr(0);
						setState(118);
						match(T__8);
						}
						break;
					}
					} 
				}
				setState(124);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		case 1:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 29);
		case 1:
			return precpred(_ctx, 28);
		case 2:
			return precpred(_ctx, 27);
		case 3:
			return precpred(_ctx, 26);
		case 4:
			return precpred(_ctx, 25);
		case 5:
			return precpred(_ctx, 24);
		case 6:
			return precpred(_ctx, 21);
		case 7:
			return precpred(_ctx, 20);
		case 8:
			return precpred(_ctx, 19);
		case 9:
			return precpred(_ctx, 18);
		case 10:
			return precpred(_ctx, 17);
		case 11:
			return precpred(_ctx, 16);
		case 12:
			return precpred(_ctx, 15);
		case 13:
			return precpred(_ctx, 23);
		case 14:
			return precpred(_ctx, 22);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\"\u0080\4\2\t\2\4"+
		"\3\t\3\3\2\7\2\b\n\2\f\2\16\2\13\13\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\7\3\26\n\3\f\3\16\3\31\13\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\7\3#"+
		"\n\3\f\3\16\3&\13\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3A\n\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\7\3{\n\3\f\3\16\3~\13\3\3\3\2\3\4\4\2\4\2\2\2\u009c\2\t\3\2\2\2\4"+
		"@\3\2\2\2\6\b\5\4\3\2\7\6\3\2\2\2\b\13\3\2\2\2\t\7\3\2\2\2\t\n\3\2\2\2"+
		"\n\3\3\2\2\2\13\t\3\2\2\2\f\r\b\3\1\2\r\16\7\16\2\2\16A\5\4\3\20\17\20"+
		"\7\22\2\2\20A\5\4\3\17\21\22\7\b\2\2\22\27\5\4\3\2\23\24\7\23\2\2\24\26"+
		"\5\4\3\2\25\23\3\2\2\2\26\31\3\2\2\2\27\25\3\2\2\2\27\30\3\2\2\2\30\32"+
		"\3\2\2\2\31\27\3\2\2\2\32\33\7\t\2\2\33A\3\2\2\2\34\35\7\n\2\2\35\36\5"+
		"\4\3\2\36\37\7\13\2\2\37A\3\2\2\2 $\7\24\2\2!#\5\4\3\2\"!\3\2\2\2#&\3"+
		"\2\2\2$\"\3\2\2\2$%\3\2\2\2%\'\3\2\2\2&$\3\2\2\2\'A\7\25\2\2()\7\26\2"+
		"\2)*\5\4\3\2*+\7\27\2\2+,\5\4\3\2,-\7\30\2\2-.\5\4\3\13.A\3\2\2\2/\60"+
		"\7\31\2\2\60\61\5\4\3\2\61\62\7\7\2\2\62\63\5\4\3\2\63\64\7\32\2\2\64"+
		"\65\5\4\3\n\65A\3\2\2\2\66\67\7\33\2\2\67A\5\4\3\t89\7\34\2\29A\5\4\3"+
		"\b:;\7\35\2\2;A\5\4\3\7<A\7\36\2\2=A\7\37\2\2>A\7 \2\2?A\7!\2\2@\f\3\2"+
		"\2\2@\17\3\2\2\2@\21\3\2\2\2@\34\3\2\2\2@ \3\2\2\2@(\3\2\2\2@/\3\2\2\2"+
		"@\66\3\2\2\2@8\3\2\2\2@:\3\2\2\2@<\3\2\2\2@=\3\2\2\2@>\3\2\2\2@?\3\2\2"+
		"\2A|\3\2\2\2BC\f\37\2\2CD\7\3\2\2D{\5\4\3 EF\f\36\2\2FG\7\4\2\2G{\5\4"+
		"\3\37HI\f\35\2\2IJ\7\5\2\2J{\5\4\3\36KL\f\34\2\2LM\7\6\2\2MN\5\4\3\2N"+
		"O\7\7\2\2OP\5\4\3\35P{\3\2\2\2QR\f\33\2\2RS\7\6\2\2S{\5\4\3\34TU\f\32"+
		"\2\2UV\7\b\2\2VW\5\4\3\2WX\7\t\2\2XY\7\7\2\2YZ\5\4\3\33Z{\3\2\2\2[\\\f"+
		"\27\2\2\\]\7\f\2\2]{\5\4\3\30^_\f\26\2\2_`\7\r\2\2`{\5\4\3\27ab\f\25\2"+
		"\2bc\7\16\2\2c{\5\4\3\26de\f\24\2\2ef\7\17\2\2f{\5\4\3\25gh\f\23\2\2h"+
		"i\7\20\2\2i{\5\4\3\24jk\f\22\2\2kl\7\21\2\2l{\5\4\3\23mn\f\21\2\2no\7"+
		"\7\2\2o{\5\4\3\22pq\f\31\2\2qr\7\b\2\2rs\5\4\3\2st\7\t\2\2t{\3\2\2\2u"+
		"v\f\30\2\2vw\7\n\2\2wx\5\4\3\2xy\7\13\2\2y{\3\2\2\2zB\3\2\2\2zE\3\2\2"+
		"\2zH\3\2\2\2zK\3\2\2\2zQ\3\2\2\2zT\3\2\2\2z[\3\2\2\2z^\3\2\2\2za\3\2\2"+
		"\2zd\3\2\2\2zg\3\2\2\2zj\3\2\2\2zm\3\2\2\2zp\3\2\2\2zu\3\2\2\2{~\3\2\2"+
		"\2|z\3\2\2\2|}\3\2\2\2}\5\3\2\2\2~|\3\2\2\2\b\t\27$@z|";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
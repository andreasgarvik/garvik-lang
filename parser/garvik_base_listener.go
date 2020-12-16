// Code generated from Garvik.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Garvik

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGarvikListener is a complete listener for a parse tree produced by GarvikParser.
type BaseGarvikListener struct{}

var _ GarvikListener = &BaseGarvikListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGarvikListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGarvikListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGarvikListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGarvikListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGarvikListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGarvikListener) ExitProgram(ctx *ProgramContext) {}

// EnterStrExpr is called when production strExpr is entered.
func (s *BaseGarvikListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production strExpr is exited.
func (s *BaseGarvikListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterDotExpr is called when production dotExpr is entered.
func (s *BaseGarvikListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production dotExpr is exited.
func (s *BaseGarvikListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterLenExpr is called when production lenExpr is entered.
func (s *BaseGarvikListener) EnterLenExpr(ctx *LenExprContext) {}

// ExitLenExpr is called when production lenExpr is exited.
func (s *BaseGarvikListener) ExitLenExpr(ctx *LenExprContext) {}

// EnterNegativeExpr is called when production negativeExpr is entered.
func (s *BaseGarvikListener) EnterNegativeExpr(ctx *NegativeExprContext) {}

// ExitNegativeExpr is called when production negativeExpr is exited.
func (s *BaseGarvikListener) ExitNegativeExpr(ctx *NegativeExprContext) {}

// EnterPopExpr is called when production popExpr is entered.
func (s *BaseGarvikListener) EnterPopExpr(ctx *PopExprContext) {}

// ExitPopExpr is called when production popExpr is exited.
func (s *BaseGarvikListener) ExitPopExpr(ctx *PopExprContext) {}

// EnterSubExpr is called when production subExpr is entered.
func (s *BaseGarvikListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production subExpr is exited.
func (s *BaseGarvikListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseGarvikListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseGarvikListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterLetExpr is called when production letExpr is entered.
func (s *BaseGarvikListener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production letExpr is exited.
func (s *BaseGarvikListener) ExitLetExpr(ctx *LetExprContext) {}

// EnterGreaterExpr is called when production greaterExpr is entered.
func (s *BaseGarvikListener) EnterGreaterExpr(ctx *GreaterExprContext) {}

// ExitGreaterExpr is called when production greaterExpr is exited.
func (s *BaseGarvikListener) ExitGreaterExpr(ctx *GreaterExprContext) {}

// EnterLambdaExpr is called when production lambdaExpr is entered.
func (s *BaseGarvikListener) EnterLambdaExpr(ctx *LambdaExprContext) {}

// ExitLambdaExpr is called when production lambdaExpr is exited.
func (s *BaseGarvikListener) ExitLambdaExpr(ctx *LambdaExprContext) {}

// EnterStructExpr is called when production structExpr is entered.
func (s *BaseGarvikListener) EnterStructExpr(ctx *StructExprContext) {}

// ExitStructExpr is called when production structExpr is exited.
func (s *BaseGarvikListener) ExitStructExpr(ctx *StructExprContext) {}

// EnterDivExpr is called when production divExpr is entered.
func (s *BaseGarvikListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production divExpr is exited.
func (s *BaseGarvikListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterLookupAssignExpr is called when production lookupAssignExpr is entered.
func (s *BaseGarvikListener) EnterLookupAssignExpr(ctx *LookupAssignExprContext) {}

// ExitLookupAssignExpr is called when production lookupAssignExpr is exited.
func (s *BaseGarvikListener) ExitLookupAssignExpr(ctx *LookupAssignExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseGarvikListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseGarvikListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterListExpr is called when production listExpr is entered.
func (s *BaseGarvikListener) EnterListExpr(ctx *ListExprContext) {}

// ExitListExpr is called when production listExpr is exited.
func (s *BaseGarvikListener) ExitListExpr(ctx *ListExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseGarvikListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseGarvikListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterIfElseExpr is called when production ifElseExpr is entered.
func (s *BaseGarvikListener) EnterIfElseExpr(ctx *IfElseExprContext) {}

// ExitIfElseExpr is called when production ifElseExpr is exited.
func (s *BaseGarvikListener) ExitIfElseExpr(ctx *IfElseExprContext) {}

// EnterLessExpr is called when production lessExpr is entered.
func (s *BaseGarvikListener) EnterLessExpr(ctx *LessExprContext) {}

// ExitLessExpr is called when production lessExpr is exited.
func (s *BaseGarvikListener) ExitLessExpr(ctx *LessExprContext) {}

// EnterDotAssignExpr is called when production dotAssignExpr is entered.
func (s *BaseGarvikListener) EnterDotAssignExpr(ctx *DotAssignExprContext) {}

// ExitDotAssignExpr is called when production dotAssignExpr is exited.
func (s *BaseGarvikListener) ExitDotAssignExpr(ctx *DotAssignExprContext) {}

// EnterMultExpr is called when production multExpr is entered.
func (s *BaseGarvikListener) EnterMultExpr(ctx *MultExprContext) {}

// ExitMultExpr is called when production multExpr is exited.
func (s *BaseGarvikListener) ExitMultExpr(ctx *MultExprContext) {}

// EnterFieldExpr is called when production fieldExpr is entered.
func (s *BaseGarvikListener) EnterFieldExpr(ctx *FieldExprContext) {}

// ExitFieldExpr is called when production fieldExpr is exited.
func (s *BaseGarvikListener) ExitFieldExpr(ctx *FieldExprContext) {}

// EnterCommentExpr is called when production commentExpr is entered.
func (s *BaseGarvikListener) EnterCommentExpr(ctx *CommentExprContext) {}

// ExitCommentExpr is called when production commentExpr is exited.
func (s *BaseGarvikListener) ExitCommentExpr(ctx *CommentExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseGarvikListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseGarvikListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterDropExpr is called when production dropExpr is entered.
func (s *BaseGarvikListener) EnterDropExpr(ctx *DropExprContext) {}

// ExitDropExpr is called when production dropExpr is exited.
func (s *BaseGarvikListener) ExitDropExpr(ctx *DropExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseGarvikListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseGarvikListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterLookupExpr is called when production lookupExpr is entered.
func (s *BaseGarvikListener) EnterLookupExpr(ctx *LookupExprContext) {}

// ExitLookupExpr is called when production lookupExpr is exited.
func (s *BaseGarvikListener) ExitLookupExpr(ctx *LookupExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseGarvikListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseGarvikListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIdExpr is called when production idExpr is entered.
func (s *BaseGarvikListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production idExpr is exited.
func (s *BaseGarvikListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterEqualExpr is called when production equalExpr is entered.
func (s *BaseGarvikListener) EnterEqualExpr(ctx *EqualExprContext) {}

// ExitEqualExpr is called when production equalExpr is exited.
func (s *BaseGarvikListener) ExitEqualExpr(ctx *EqualExprContext) {}

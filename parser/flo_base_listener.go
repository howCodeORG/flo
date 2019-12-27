// Generated from Flo.g4 by ANTLR 4.7.

package parser // Flo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFloListener is a complete listener for a parse tree produced by FloParser.
type BaseFloListener struct{}

var _ FloListener = &BaseFloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseFloListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseFloListener) ExitStart(ctx *StartContext) {}

// EnterMulti_stmts is called when production multi_stmts is entered.
func (s *BaseFloListener) EnterMulti_stmts(ctx *Multi_stmtsContext) {}

// ExitMulti_stmts is called when production multi_stmts is exited.
func (s *BaseFloListener) ExitMulti_stmts(ctx *Multi_stmtsContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseFloListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseFloListener) ExitStmt(ctx *StmtContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BaseFloListener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BaseFloListener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseFloListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseFloListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseFloListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseFloListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseFloListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseFloListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterFunc_decl is called when production func_decl is entered.
func (s *BaseFloListener) EnterFunc_decl(ctx *Func_declContext) {}

// ExitFunc_decl is called when production func_decl is exited.
func (s *BaseFloListener) ExitFunc_decl(ctx *Func_declContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseFloListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseFloListener) ExitParameters(ctx *ParametersContext) {}

// EnterFor_clause_block is called when production for_clause_block is entered.
func (s *BaseFloListener) EnterFor_clause_block(ctx *For_clause_blockContext) {}

// ExitFor_clause_block is called when production for_clause_block is exited.
func (s *BaseFloListener) ExitFor_clause_block(ctx *For_clause_blockContext) {}

// EnterCondition_block is called when production condition_block is entered.
func (s *BaseFloListener) EnterCondition_block(ctx *Condition_blockContext) {}

// ExitCondition_block is called when production condition_block is exited.
func (s *BaseFloListener) ExitCondition_block(ctx *Condition_blockContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BaseFloListener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BaseFloListener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseFloListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseFloListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseFloListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseFloListener) ExitBlock(ctx *BlockContext) {}

// EnterUnaryIncDec is called when production UnaryIncDec is entered.
func (s *BaseFloListener) EnterUnaryIncDec(ctx *UnaryIncDecContext) {}

// ExitUnaryIncDec is called when production UnaryIncDec is exited.
func (s *BaseFloListener) ExitUnaryIncDec(ctx *UnaryIncDecContext) {}

// EnterReadIdentifier is called when production ReadIdentifier is entered.
func (s *BaseFloListener) EnterReadIdentifier(ctx *ReadIdentifierContext) {}

// ExitReadIdentifier is called when production ReadIdentifier is exited.
func (s *BaseFloListener) ExitReadIdentifier(ctx *ReadIdentifierContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseFloListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseFloListener) ExitOr(ctx *OrContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseFloListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseFloListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseFloListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseFloListener) ExitAddSub(ctx *AddSubContext) {}

// EnterBitShift is called when production BitShift is entered.
func (s *BaseFloListener) EnterBitShift(ctx *BitShiftContext) {}

// ExitBitShift is called when production BitShift is exited.
func (s *BaseFloListener) ExitBitShift(ctx *BitShiftContext) {}

// EnterExpressionGroup is called when production ExpressionGroup is entered.
func (s *BaseFloListener) EnterExpressionGroup(ctx *ExpressionGroupContext) {}

// ExitExpressionGroup is called when production ExpressionGroup is exited.
func (s *BaseFloListener) ExitExpressionGroup(ctx *ExpressionGroupContext) {}

// EnterRelational is called when production Relational is entered.
func (s *BaseFloListener) EnterRelational(ctx *RelationalContext) {}

// ExitRelational is called when production Relational is exited.
func (s *BaseFloListener) ExitRelational(ctx *RelationalContext) {}

// EnterString is called when production String is entered.
func (s *BaseFloListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseFloListener) ExitString(ctx *StringContext) {}

// EnterUnary is called when production Unary is entered.
func (s *BaseFloListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production Unary is exited.
func (s *BaseFloListener) ExitUnary(ctx *UnaryContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseFloListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseFloListener) ExitInteger(ctx *IntegerContext) {}

// EnterGetItem is called when production GetItem is entered.
func (s *BaseFloListener) EnterGetItem(ctx *GetItemContext) {}

// ExitGetItem is called when production GetItem is exited.
func (s *BaseFloListener) ExitGetItem(ctx *GetItemContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseFloListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseFloListener) ExitFloat(ctx *FloatContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseFloListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseFloListener) ExitNot(ctx *NotContext) {}

// EnterBitwiseOr is called when production BitwiseOr is entered.
func (s *BaseFloListener) EnterBitwiseOr(ctx *BitwiseOrContext) {}

// ExitBitwiseOr is called when production BitwiseOr is exited.
func (s *BaseFloListener) ExitBitwiseOr(ctx *BitwiseOrContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseFloListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseFloListener) ExitAnd(ctx *AndContext) {}

// EnterBitwiseAnd is called when production BitwiseAnd is entered.
func (s *BaseFloListener) EnterBitwiseAnd(ctx *BitwiseAndContext) {}

// ExitBitwiseAnd is called when production BitwiseAnd is exited.
func (s *BaseFloListener) ExitBitwiseAnd(ctx *BitwiseAndContext) {}

// EnterList is called when production List is entered.
func (s *BaseFloListener) EnterList(ctx *ListContext) {}

// ExitList is called when production List is exited.
func (s *BaseFloListener) ExitList(ctx *ListContext) {}

// EnterXOR is called when production XOR is entered.
func (s *BaseFloListener) EnterXOR(ctx *XORContext) {}

// ExitXOR is called when production XOR is exited.
func (s *BaseFloListener) ExitXOR(ctx *XORContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseFloListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseFloListener) ExitBoolean(ctx *BooleanContext) {}

// EnterCallExpression is called when production CallExpression is entered.
func (s *BaseFloListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production CallExpression is exited.
func (s *BaseFloListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterPower is called when production Power is entered.
func (s *BaseFloListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production Power is exited.
func (s *BaseFloListener) ExitPower(ctx *PowerContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseFloListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseFloListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseFloListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseFloListener) ExitEos(ctx *EosContext) {}

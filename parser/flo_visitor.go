// Generated from Flo.g4 by ANTLR 4.7.

package parser // Flo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FloParser.
type FloVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FloParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by FloParser#multi_stmts.
	VisitMulti_stmts(ctx *Multi_stmtsContext) interface{}

	// Visit a parse tree produced by FloParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by FloParser#simple_stmt.
	VisitSimple_stmt(ctx *Simple_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#func_decl.
	VisitFunc_decl(ctx *Func_declContext) interface{}

	// Visit a parse tree produced by FloParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by FloParser#for_clause_block.
	VisitFor_clause_block(ctx *For_clause_blockContext) interface{}

	// Visit a parse tree produced by FloParser#condition_block.
	VisitCondition_block(ctx *Condition_blockContext) interface{}

	// Visit a parse tree produced by FloParser#print_stmt.
	VisitPrint_stmt(ctx *Print_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by FloParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by FloParser#UnaryIncDec.
	VisitUnaryIncDec(ctx *UnaryIncDecContext) interface{}

	// Visit a parse tree produced by FloParser#ReadIdentifier.
	VisitReadIdentifier(ctx *ReadIdentifierContext) interface{}

	// Visit a parse tree produced by FloParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by FloParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by FloParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by FloParser#BitShift.
	VisitBitShift(ctx *BitShiftContext) interface{}

	// Visit a parse tree produced by FloParser#ExpressionGroup.
	VisitExpressionGroup(ctx *ExpressionGroupContext) interface{}

	// Visit a parse tree produced by FloParser#Relational.
	VisitRelational(ctx *RelationalContext) interface{}

	// Visit a parse tree produced by FloParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by FloParser#Unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by FloParser#Integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by FloParser#GetItem.
	VisitGetItem(ctx *GetItemContext) interface{}

	// Visit a parse tree produced by FloParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by FloParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by FloParser#BitwiseOr.
	VisitBitwiseOr(ctx *BitwiseOrContext) interface{}

	// Visit a parse tree produced by FloParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by FloParser#BitwiseAnd.
	VisitBitwiseAnd(ctx *BitwiseAndContext) interface{}

	// Visit a parse tree produced by FloParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by FloParser#XOR.
	VisitXOR(ctx *XORContext) interface{}

	// Visit a parse tree produced by FloParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by FloParser#CallExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by FloParser#Power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by FloParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by FloParser#eos.
	VisitEos(ctx *EosContext) interface{}
}

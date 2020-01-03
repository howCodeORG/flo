// Generated from Flo.g4 by ANTLR 4.7.

package parser // Flo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FloListener is a complete listener for a parse tree produced by FloParser.
type FloListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterMulti_stmts is called when entering the multi_stmts production.
	EnterMulti_stmts(c *Multi_stmtsContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterFunc_decl is called when entering the func_decl production.
	EnterFunc_decl(c *Func_declContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterFor_clause_block is called when entering the for_clause_block production.
	EnterFor_clause_block(c *For_clause_blockContext)

	// EnterCondition_block is called when entering the condition_block production.
	EnterCondition_block(c *Condition_blockContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterUnaryIncDec is called when entering the UnaryIncDec production.
	EnterUnaryIncDec(c *UnaryIncDecContext)

	// EnterAnonFunc is called when entering the AnonFunc production.
	EnterAnonFunc(c *AnonFuncContext)

	// EnterReadIdentifier is called when entering the ReadIdentifier production.
	EnterReadIdentifier(c *ReadIdentifierContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterBitShift is called when entering the BitShift production.
	EnterBitShift(c *BitShiftContext)

	// EnterExpressionGroup is called when entering the ExpressionGroup production.
	EnterExpressionGroup(c *ExpressionGroupContext)

	// EnterRelational is called when entering the Relational production.
	EnterRelational(c *RelationalContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterUnary is called when entering the Unary production.
	EnterUnary(c *UnaryContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterGetItem is called when entering the GetItem production.
	EnterGetItem(c *GetItemContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterBitwiseOr is called when entering the BitwiseOr production.
	EnterBitwiseOr(c *BitwiseOrContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterBitwiseAnd is called when entering the BitwiseAnd production.
	EnterBitwiseAnd(c *BitwiseAndContext)

	// EnterList is called when entering the List production.
	EnterList(c *ListContext)

	// EnterXOR is called when entering the XOR production.
	EnterXOR(c *XORContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterCallExpression is called when entering the CallExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitMulti_stmts is called when exiting the multi_stmts production.
	ExitMulti_stmts(c *Multi_stmtsContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitFunc_decl is called when exiting the func_decl production.
	ExitFunc_decl(c *Func_declContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitFor_clause_block is called when exiting the for_clause_block production.
	ExitFor_clause_block(c *For_clause_blockContext)

	// ExitCondition_block is called when exiting the condition_block production.
	ExitCondition_block(c *Condition_blockContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitUnaryIncDec is called when exiting the UnaryIncDec production.
	ExitUnaryIncDec(c *UnaryIncDecContext)

	// ExitAnonFunc is called when exiting the AnonFunc production.
	ExitAnonFunc(c *AnonFuncContext)

	// ExitReadIdentifier is called when exiting the ReadIdentifier production.
	ExitReadIdentifier(c *ReadIdentifierContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitBitShift is called when exiting the BitShift production.
	ExitBitShift(c *BitShiftContext)

	// ExitExpressionGroup is called when exiting the ExpressionGroup production.
	ExitExpressionGroup(c *ExpressionGroupContext)

	// ExitRelational is called when exiting the Relational production.
	ExitRelational(c *RelationalContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitUnary is called when exiting the Unary production.
	ExitUnary(c *UnaryContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitGetItem is called when exiting the GetItem production.
	ExitGetItem(c *GetItemContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitBitwiseOr is called when exiting the BitwiseOr production.
	ExitBitwiseOr(c *BitwiseOrContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitBitwiseAnd is called when exiting the BitwiseAnd production.
	ExitBitwiseAnd(c *BitwiseAndContext)

	// ExitList is called when exiting the List production.
	ExitList(c *ListContext)

	// ExitXOR is called when exiting the XOR production.
	ExitXOR(c *XORContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitCallExpression is called when exiting the CallExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}

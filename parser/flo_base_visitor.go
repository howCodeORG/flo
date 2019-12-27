// Generated from Flo.g4 by ANTLR 4.7.

package parser // Flo

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFloVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFloVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitMulti_stmts(ctx *Multi_stmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitSimple_stmt(ctx *Simple_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitFunc_decl(ctx *Func_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitFor_clause_block(ctx *For_clause_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitCondition_block(ctx *Condition_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitPrint_stmt(ctx *Print_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitUnaryIncDec(ctx *UnaryIncDecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitReadIdentifier(ctx *ReadIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitBitShift(ctx *BitShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitExpressionGroup(ctx *ExpressionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitRelational(ctx *RelationalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitGetItem(ctx *GetItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitBitwiseOr(ctx *BitwiseOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitBitwiseAnd(ctx *BitwiseAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitXOR(ctx *XORContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFloVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

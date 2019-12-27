package compiler

import (
	"flo/parser"
	"flo/vm"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// FloVisitor holds the ANTLR stack
type FloVisitor struct {
	*parser.BaseFloVisitor

	Stack       []vm.FloObject
	blockScope  int
	funcScope   int
	environment []map[string]vm.FloObject
	Error       bool
	Object      vm.Object
}

func (v *FloVisitor) cleanup() {
	v.blockScope = 0
	v.funcScope = 0
}

func (v *FloVisitor) push(o vm.FloObject) {
	v.Stack = append(v.Stack, o)
}

func (v *FloVisitor) pop() vm.FloObject {

	if len(v.Stack) > 0 {
		x := v.Stack[len(v.Stack)-1]
		v.Stack = v.Stack[:len(v.Stack)-1]
		return x
	}
	return nil
}

// Init creates the environment for Flo program state
func (v *FloVisitor) Init() {
	v.environment = make([]map[string]vm.FloObject, 1)
	v.environment[v.blockScope] = make(map[string]vm.FloObject, 1)
}

func (v *FloVisitor) readIdentifier(name string) vm.FloObject {

	// var scope int

	// if scope == 0 {
	// 	scope = v.blockScope
	// }
	// Lookup identifier value
	// var found bool
	// for i := v.blockScope; i >= 0; i-- {
	// 	if value, ok := v.environment[i][name]; ok {
	// 		found = true
	// 		return value
	// 	}
	// }
	// if !found {
	// 	v.doError(nameNotDefinedError(name))
	// }
	return nil
}

func (v *FloVisitor) storeIdentifier(identifier string, value vm.FloObject) {

	// var found bool
	// for i := v.blockScope; i >= v.funcScope; i-- {
	// 	if _, ok := v.environment[i][identifier]; ok {
	// 		found = true
	// 		v.environment[i][identifier] = value
	// 	}
	// }
	// if !found {
	// 	v.environment[v.blockScope][identifier] = value
	// }

	v.Object.Compile_Store_Name(vm.FloString(identifier))

}

// EnterBlock increases the current scope
func (v *FloVisitor) enterBlock() {
	v.Object.Compile_Push_Block()
}

// ExitBlock increases the current scope
func (v *FloVisitor) exitBlock() {
	v.Object.Compile_Pop_Block()
}

func (v *FloVisitor) doError(err string) {
	v.cleanup()
	v.Stack = make([]vm.FloObject, 0)
	panic(err)
}

// Visit 'visits' each node in the parse tree
func (v *FloVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

// VisitStart visits the first node in the parse tree
func (v *FloVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	v.Visit(ctx.Multi_stmts())
	return nil
}

// VisitMulti_stmts evaluates multiple expressions
func (v *FloVisitor) VisitMulti_stmts(ctx *parser.Multi_stmtsContext) interface{} {
	for _, value := range ctx.AllStmt() {

		v.Visit(value)

	}

	return nil
}

// VisitStmt evaluates multiple expressions
func (v *FloVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {

	if ctx.Simple_stmt() != nil {
		// Simple_stmt
		v.Visit(ctx.Simple_stmt())
	} else if ctx.If_stmt() != nil {
		// If_stmt
		v.Visit(ctx.If_stmt())
	} else if ctx.For_stmt() != nil {
		// For_stmt
		v.Visit(ctx.For_stmt())
	} else if ctx.Func_decl() != nil {
		// Func_decl
		v.Visit(ctx.Func_decl())
	} else if ctx.Print_stmt() != nil {
		// Print_stmt
		v.Visit(ctx.Print_stmt())
	} else if ctx.Return_stmt() != nil {
		// Return_stmt
		v.Visit(ctx.Return_stmt())
	} else if ctx.Block() != nil {
		// Block
		v.Visit(ctx.Block())
	} else {
		panic("Unimplemented statement")
	}

	return nil
}

// VisitSimple_stmt evaluates simple statements
func (v *FloVisitor) VisitSimple_stmt(ctx *parser.Simple_stmtContext) interface{} {

	if ctx.Expression() != nil {
		// Expression
		v.Visit(ctx.Expression())
	} else if ctx.Assign_stmt() != nil {
		// Assign_stmt
		v.Visit(ctx.Assign_stmt())
	}

	return nil
}

// VisitExpression is used to visit the appropriate expression
func (v *FloVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {

	ctx.Accept(v)

	return nil

}

// VisitAssign_stmt evaluates Flo identifiers
func (v *FloVisitor) VisitAssign_stmt(ctx *parser.Assign_stmtContext) interface{} {

	identifier := ctx.GetStart().GetText()
	v.Visit(ctx.Expression())

	v.Object.Compile_Store_Name(vm.FloString(identifier))

	return nil

}

// VisitIf_stmt evaluates Flo if statements
func (v *FloVisitor) VisitIf_stmt(ctx *parser.If_stmtContext) interface{} {

	// If and elif's
	nops := []int{}
	for _, value := range ctx.AllCondition_block() {

		v.Visit(value)
		nops = append(nops, len(v.Object.Instructions))
		v.Object.Compile_Nop(1)

	}

	// Found else block
	if ctx.Multi_stmts() != nil {
		v.Visit(ctx.Multi_stmts())
	}
	jump_by := len(v.Object.Instructions)

	for i, _ := range nops {
		v.Object.Compile_Jump(nops[i], jump_by-nops[i])
	}

	return nil

}

// VisitCondition_block evaluates conditions
func (v *FloVisitor) VisitCondition_block(ctx *parser.Condition_blockContext) interface{} {

	v.Visit(ctx.Expression())

	start := len(v.Object.Instructions)
	v.Object.Compile_Nop(1)
	v.Visit(ctx.Multi_stmts())

	jump_by := len(v.Object.Instructions) - start + 2
	v.Object.Compile_Pop_Jump_If_False(start, jump_by)

	return nil
}

// VisitFor_stmt evaluates Flo for statements/loops
func (v *FloVisitor) VisitFor_stmt(ctx *parser.For_stmtContext) interface{} {

	v.enterBlock()

	// For loop with for-clause
	if ctx.For_clause_block() != nil {

		v.Visit(ctx.For_clause_block())

	}

	v.exitBlock()

	return nil

}

// VisitFor_clause_block evaluates Flo for loop clauses and the block
func (v *FloVisitor) VisitFor_clause_block(ctx *parser.For_clause_blockContext) interface{} {

	// Initial var assignment
	v.Visit(ctx.Simple_stmt(0))
	start := len(v.Object.Instructions)
	// v.Object.Compile_Nop(1)
	// Condition
	v.Visit(ctx.Expression())
	block_start := len(v.Object.Instructions)
	v.Object.Compile_Nop(1)
	// Execute block
	v.Visit(ctx.Multi_stmts())
	// Increase counter
	v.Visit(ctx.Simple_stmt(1))
	block_end := len(v.Object.Instructions)
	v.Object.Compile_Pop_Jump_If_False(block_start, block_end-block_start+2)
	v.Object.Compile_Jump_Back(block_end - start)
	return nil

}

// VisitFunc_decl evaluates Flo function definitions
func (v *FloVisitor) VisitFunc_decl(ctx *parser.Func_declContext) interface{} {

	name := ctx.IDENTIFIER()
	// Setup function sets gets the functions name.
	v.Object.Compile_Setup_Function(vm.FloString(name.GetText()))

	// Setup parameters gets function parameters
	params := v.Visit(ctx.Parameters())
	if params == nil {
		v.Object.Compile_Setup_Params(0)
	} else {
		v.Object.Compile_Setup_Params(params.(vm.FloInteger))
	}

	// Setup body collects all instructions that make up the function body
	v.Object.Compile_Setup_Body()
	v.Object.Compile_Push_Block()
	v.Visit(ctx.Multi_stmts())
	v.Object.Compile_Pop_Block()

	// Make function uses the output of the previous instructions to create a FloCallable object.
	v.Object.Compile_Make_Function()

	// Store function in variable
	v.Object.Compile_Store_Name(vm.FloString(name.GetText()))

	// v.Object.Compile_Store_Name(vm.FloString(name.GetText()))
	// val := FloCallable{
	// 	name:      name.GetText(),
	// 	min_scope: v.blockScope,
	// 	args:      args,
	// 	body:      ctx.Block(),
	// }

	// v.storeIdentifier(name.GetText(), val)

	return nil

}

// VisitParameters evaluates Flo callable parameters
func (v *FloVisitor) VisitParameters(ctx *parser.ParametersContext) interface{} {

	vals := make([]vm.FloObject, 0)
	for _, val := range ctx.AllIDENTIFIER() {
		// vals = append(vals, v.Visit(val).(FloString))
		vals = append(vals, vm.FloString(val.GetText()))

		v.Object.Compile_Load_Const(vm.FloString(val.GetText()))
	}
	if vals == nil {
		return vm.FloInteger(0)
	}
	return vm.FloInteger(len(vals))
	// return nil
}

// VisitPrint_stmt executes 'print'
func (v *FloVisitor) VisitPrint_stmt(ctx *parser.Print_stmtContext) interface{} {
	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
		// fmt.Println(x)
	}
	v.Object.Compile_Print()
	return nil
}

// VisitExpressionGroup is used to evaluate expression groups
func (v *FloVisitor) VisitExpressionGroup(ctx *parser.ExpressionGroupContext) interface{} {

	return ctx.Expression().Accept(v)
}

// VisitMulDiv evaluates multiplication and division nodes in the parse tree
func (v *FloVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserMULTIPLY:
		v.Object.Compile_Binary_Mul()
		return nil
	case parser.FloParserDIVIDE:
		v.Object.Compile_Binary_Div()
		return nil
	case parser.FloParserMOD:
		v.Object.Compile_Binary_Mod()
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitAddSub evaluates addition and subtraction nodes in the parse tree
func (v *FloVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserADDITION:
		v.Object.Compile_Binary_Add()
		return nil
	case parser.FloParserSUBTRACTION:
		v.Object.Compile_Binary_Sub()
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitBitShift evaluates << and >> nodes in the parse tree
func (v *FloVisitor) VisitBitShift(ctx *parser.BitShiftContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserLSHIFT:
		v.Object.Compile_Binary_LShift()
		return nil
	case parser.FloParserRSHIFT:
		v.Object.Compile_Binary_RShift()
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitRelational evaluates <, <=, >, >=, ==, !=
func (v *FloVisitor) VisitRelational(ctx *parser.RelationalContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserLESS:
		v.Object.Compile_Compare(0)
		return nil
	case parser.FloParserLESS_OR_EQUALS:
		v.Object.Compile_Compare(1)
		return nil
	case parser.FloParserGREATER:
		v.Object.Compile_Compare(2)
		return nil
	case parser.FloParserGREATER_OR_EQUALS:
		v.Object.Compile_Compare(3)
		return nil
	case parser.FloParserEQUALS:
		v.Object.Compile_Compare(4)
		return nil
	case parser.FloParserNOT_EQUALS:
		v.Object.Compile_Compare(5)
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitPower evaluates **
func (v *FloVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))
	v.Object.Compile_Binary_Pow()
	return nil
}

// VisitUnary evaluates unary +, -
func (v *FloVisitor) VisitUnary(ctx *parser.UnaryContext) interface{} {

	v.Visit(ctx.Expression())

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserADDITION:
		v.Object.Compile_Unary_Add()
		return nil
	case parser.FloParserSUBTRACTION:
		v.Object.Compile_Unary_Sub()
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitUnaryIncDec evaluates unary ++, --
func (v *FloVisitor) VisitUnaryIncDec(ctx *parser.UnaryIncDecContext) interface{} {

	// value := v.Visit(ctx.Expression()).(FloObject)
	name := ctx.GetStart().GetText()
	// val := v.readIdentifier(name)

	switch ctx.GetOp().GetTokenType() {
	case parser.FloParserPLUS_PLUS:
		v.Object.Compile_Unary_Inc(vm.FloString(name))
		// new_val := UnaryPlusPlus(val)
		// v.storeIdentifier(name, new_val)
		return nil
	case parser.FloParserMINUS_MINUS:
		v.Object.Compile_Unary_Dec(vm.FloString(name))
		// new_val := UnaryMinusMinus(val)
		// v.storeIdentifier(name, new_val)
		return nil
	default:
		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
		return nil
	}
}

// VisitBitwiseAnd evaluates &
func (v *FloVisitor) VisitBitwiseAnd(ctx *parser.BitwiseAndContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	v.Object.Compile_Binary_And()
	return nil
}

// VisitXOR evaluates ^
func (v *FloVisitor) VisitXOR(ctx *parser.XORContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	v.Object.Compile_Binary_XOR()

	return nil
}

// VisitBitwiseOr evaluates |
func (v *FloVisitor) VisitBitwiseOr(ctx *parser.BitwiseOrContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	v.Object.Compile_Binary_Or()
	return nil
}

// VisitAnd evaluates logical 'and'
func (v *FloVisitor) VisitAnd(ctx *parser.AndContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))
	v.Object.Compile_Logical_And()
	return nil
}

// VisitOr evaluates logical 'or'
func (v *FloVisitor) VisitOr(ctx *parser.OrContext) interface{} {

	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))
	v.Object.Compile_Logical_Or()
	return nil
}

// VisitNot evaluates logical 'not'
func (v *FloVisitor) VisitNot(ctx *parser.NotContext) interface{} {

	v.Visit(ctx.Expression())
	v.Object.Compile_Unary_Not()
	// x := Not(value)
	return nil
}

// VisitList evaluates Flo lists
func (v *FloVisitor) VisitList(ctx *parser.ListContext) interface{} {

	list := make([]vm.FloObject, 0)

	if ctx.Expression_list() != nil {
		vals := v.Visit(ctx.Expression_list()).(vm.FloList)
		list = append(list, vals...)
	}

	v.Object.Compile_Build_List(uint32(len(list)))

	return nil

}

// VisitExpression_list evaluates Flo expression lists
func (v *FloVisitor) VisitExpression_list(ctx *parser.Expression_listContext) interface{} {

	vals := make([]vm.FloObject, 0)
	for _, val := range ctx.AllExpression() {
		v.Visit(val)
		vals = append(vals, vm.FloInteger(0))

	}

	return vm.FloList(vals)
}

// VisitGetItem evaluates Flo list indicies
func (v *FloVisitor) VisitGetItem(ctx *parser.GetItemContext) interface{} {
	// lval
	v.Visit(ctx.Expression(0))

	// switch lval.(type) {
	// case FloList:
	// 	break
	// default:
	// 	v.doError(listTypeError(lval.(FloObject)))
	// }
	// Index
	v.Visit(ctx.Expression(1))
	// switch index.(type) {
	// case FloInteger:
	// 	break
	// default:
	// 	v.doError(listIndexTypeError(index.(FloObject)))
	// }

	v.Object.Compile_Get_Item()

	// list := lval.(FloList)
	// list_index := index.(FloInteger)

	// if int(list_index) >= len(list) || int(list_index) < 0 {
	// 	v.doError(indexOutOfRangeError())
	// }

	// return list[list_index]
	return nil
}

func (v *FloVisitor) callCallable(fn vm.FloCallable, args vm.FloList) vm.FloObject {

	// prev_func_scope := v.funcScope

	// v.push(fn)

	// v.enterBlock()

	// v.funcScope = v.blockScope

	// // Check arity
	// if len(args) != len(fn.Args) {
	// 	v.doError(vm.allableArgsTypeError(fn, len(args)))
	// }

	// // Initialise callable with parameters
	// for i, arg := range args {
	// 	v.storeIdentifier(string(fn.args[i].(FloString)), arg)
	// }

	// // Call function
	// fn.body.Accept(v)

	// // Pop return value off call stack
	// x := v.pop()

	// // Pop function off call stack
	// v.pop()

	// v.exitBlock()

	// v.funcScope = prev_func_scope

	// if x == nil {
	// 	return nil
	// }
	// // fmt.Println(v.Stack)
	// return x.(FloObject)
	return nil
}

// VisitCallExpression evaluates Flo function calls
func (v *FloVisitor) VisitCallExpression(ctx *parser.CallExpressionContext) interface{} {

	// if ctx.Expression() == nil {
	// 	v.doError("n")
	// }

	args := v.Visit(ctx.Expression_list())
	v.Visit(ctx.Expression())
	if args == nil {
		v.Object.Compile_Call_Function(0)
	} else {
		v.Object.Compile_Call_Function(len(args.(vm.FloList)))
	}

	// switch x.(type) {
	// case FloCallable:
	// 	break
	// default:
	// 	v.doError(notCallableError(x.(FloObject)))
	// }

	// fn := x.(FloCallable)

	// var args FloList
	// switch y.(type) {
	// case FloList:
	// 	args = y.(FloList)
	// 	break
	// case nil:
	// 	break
	// default:
	// 	v.doError(notCallableError(y.(FloObject)))
	// }

	// return v.callCallable(fn, args)
	return nil

}

// VisitReturn_stmt evaluates Flo function calls
func (v *FloVisitor) VisitReturn_stmt(ctx *parser.Return_stmtContext) interface{} {

	if ctx.Expression() == nil {
		return nil
	}

	v.Visit(ctx.Expression())
	v.Object.Compile_Return()
	// v.push(x.(FloObject))
	return nil

}

// VisitInteger evaluates Flo integers
func (v *FloVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {

	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}

	x := vm.FloInteger(i)

	v.Object.Compile_Load_Const(x)

	return nil

}

// VisitFloat evaluates floats in Flo
func (v *FloVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {

	f, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}

	x := vm.FloFloat(f)
	v.Object.Compile_Load_Const(x)

	return nil
}

// VisitString evaluates Flo strings
func (v *FloVisitor) VisitString(ctx *parser.StringContext) interface{} {

	value := ctx.GetText()
	value = value[1 : len(value)-1]

	x := vm.FloString(value)
	v.Object.Compile_Load_Const(x)

	return nil

}

// VisitBoolean evaluates Flo booleans
func (v *FloVisitor) VisitBoolean(ctx *parser.BooleanContext) interface{} {

	x := ctx.GetText()
	if x == "true" {
		v.Object.Compile_Load_Const(vm.FloBool(true))
	} else {
		v.Object.Compile_Load_Const(vm.FloBool(false))
	}
	return nil

}

// VisitReadIdentifier evaluates Flo identifiers
func (v *FloVisitor) VisitReadIdentifier(ctx *parser.ReadIdentifierContext) interface{} {

	identifier := ctx.GetText()

	v.Object.Compile_Load_Name(vm.FloString(identifier))
	// val := v.readIdentifier(identifier)
	// if val != nil {
	// 	return val
	// }
	return nil

}

// VisitBlock evaluates blocks
func (v *FloVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	v.enterBlock()
	if ctx.Multi_stmts() != nil {
		v.Visit(ctx.Multi_stmts())
	}
	v.exitBlock()

	return nil
}

// Visit a parse tree produced by FloParser#eos.
func (v *FloVisitor) VisitEos(ctx *parser.EosContext) interface{} {
	return nil
}

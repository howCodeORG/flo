package eval

// import (
// 	"flo/parser"
// 	"fmt"
// 	"strconv"

// 	"github.com/antlr/antlr4/runtime/Go/antlr"
// )

// // FloVisitor holds the ANTLR stack
// type FloVisitor struct {
// 	*parser.BaseFloVisitor

// 	Stack       []FloObject
// 	blockScope  int
// 	funcScope   int
// 	environment []map[string]FloObject
// 	Error       bool
// }

// func (v *FloVisitor) cleanup() {
// 	v.blockScope = 0
// 	v.funcScope = 0
// }

// func (v *FloVisitor) push(o FloObject) {
// 	v.Stack = append(v.Stack, o)
// }

// func (v *FloVisitor) pop() FloObject {

// 	if len(v.Stack) > 0 {
// 		x := v.Stack[len(v.Stack)-1]
// 		v.Stack = v.Stack[:len(v.Stack)-1]
// 		return x
// 	}
// 	return nil
// }

// // Init creates the environment for Flo program state
// func (v *FloVisitor) Init() {
// 	v.environment = make([]map[string]FloObject, 1)
// 	v.environment[v.blockScope] = make(map[string]FloObject, 1)
// }

// func (v *FloVisitor) readIdentifier(name string) FloObject {

// 	// var scope int

// 	// if scope == 0 {
// 	// 	scope = v.blockScope
// 	// }
// 	// Lookup identifier value
// 	var found bool
// 	for i := v.blockScope; i >= 0; i-- {
// 		if value, ok := v.environment[i][name]; ok {
// 			found = true
// 			return value
// 		}
// 	}
// 	if !found {
// 		v.doError(nameNotDefinedError(name))
// 	}
// 	return nil
// }

// func (v *FloVisitor) storeIdentifier(identifier string, value FloObject) {

// 	var found bool
// 	for i := v.blockScope; i >= v.funcScope; i-- {
// 		if _, ok := v.environment[i][identifier]; ok {
// 			found = true
// 			v.environment[i][identifier] = value
// 		}
// 	}
// 	if !found {
// 		v.environment[v.blockScope][identifier] = value
// 	}

// }

// // EnterBlock increases the current scope
// func (v *FloVisitor) enterBlock() {
// 	// Increase scope
// 	v.blockScope++
// 	v.environment = append(v.environment, make(map[string]FloObject, 1))
// }

// // ExitBlock increases the current scope
// func (v *FloVisitor) exitBlock() {
// 	// Decrease scope
// 	v.environment = v.environment[:v.blockScope]
// 	v.blockScope--
// }

// func (v *FloVisitor) doError(err string) {
// 	v.cleanup()
// 	v.Stack = make([]FloObject, 0)
// 	panic(err)
// }

// // Visit 'visits' each node in the parse tree
// func (v *FloVisitor) Visit(tree antlr.ParseTree) interface{} {
// 	if tree == nil {
// 		return nil
// 	}
// 	return tree.Accept(v)
// }

// // VisitStart visits the first node in the parse tree
// func (v *FloVisitor) VisitStart(ctx *parser.StartContext) interface{} {
// 	v.Visit(ctx.Multi_stmts())
// 	return nil
// }

// // VisitMulti_stmts evaluates multiple expressions
// func (v *FloVisitor) VisitMulti_stmts(ctx *parser.Multi_stmtsContext) interface{} {
// 	for _, value := range ctx.AllStmt() {

// 		v.Visit(value)

// 	}

// 	return nil
// }

// // VisitStmt evaluates multiple expressions
// func (v *FloVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {

// 	if ctx.Simple_stmt() != nil {
// 		// Simple_stmt
// 		v.Visit(ctx.Simple_stmt())
// 	} else if ctx.If_stmt() != nil {
// 		// If_stmt
// 		v.Visit(ctx.If_stmt())
// 	} else if ctx.For_stmt() != nil {
// 		// For_stmt
// 		v.Visit(ctx.For_stmt())
// 	} else if ctx.Func_decl() != nil {
// 		// Func_decl
// 		v.Visit(ctx.Func_decl())
// 	} else if ctx.Print_stmt() != nil {
// 		// Print_stmt
// 		v.Visit(ctx.Print_stmt())
// 	} else if ctx.Return_stmt() != nil {
// 		// Return_stmt
// 		v.Visit(ctx.Return_stmt())
// 	} else if ctx.Block() != nil {
// 		// Block
// 		v.Visit(ctx.Block())
// 	} else {
// 		panic("Unimplemented statement")
// 	}

// 	return nil
// }

// // VisitSimple_stmt evaluates simple statements
// func (v *FloVisitor) VisitSimple_stmt(ctx *parser.Simple_stmtContext) interface{} {

// 	if ctx.Expression() != nil {
// 		// Expression
// 		v.Visit(ctx.Expression())
// 	} else if ctx.Assign_stmt() != nil {
// 		// Assign_stmt
// 		v.Visit(ctx.Assign_stmt())
// 	}

// 	return nil
// }

// // VisitExpression is used to visit the appropriate expression
// func (v *FloVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {

// 	ctx.Accept(v)

// 	return nil

// }

// // VisitAssign_stmt evaluates Flo identifiers
// func (v *FloVisitor) VisitAssign_stmt(ctx *parser.Assign_stmtContext) interface{} {

// 	identifier := ctx.GetStart().GetText()
// 	value := v.Visit(ctx.Expression()).(FloObject)

// 	v.storeIdentifier(identifier, value)

// 	return nil

// }

// // VisitIf_stmt evaluates Flo if statements
// func (v *FloVisitor) VisitIf_stmt(ctx *parser.If_stmtContext) interface{} {

// 	v.enterBlock()
// 	var blockDone bool
// 	// If and elif's
// 	for _, value := range ctx.AllCondition_block() {

// 		condition := v.Visit(value).(FloBool)
// 		if condition {
// 			blockDone = true
// 			break
// 		}

// 	}

// 	// Found else block
// 	if ctx.Block() != nil && !blockDone {
// 		v.Visit(ctx.Block())
// 	}

// 	v.exitBlock()

// 	return nil

// }

// // VisitCondition_block evaluates conditions
// func (v *FloVisitor) VisitCondition_block(ctx *parser.Condition_blockContext) interface{} {

// 	condition := v.Visit(ctx.Expression())
// 	switch val := condition.(type) {
// 	case FloBool:
// 		if val == true {
// 			v.Visit(ctx.Block())
// 		}
// 		return val
// 	default:
// 		panic(conditionNotBoolError(condition.(FloObject)))
// 	}

// 	return false
// }

// // VisitFor_stmt evaluates Flo for statements/loops
// func (v *FloVisitor) VisitFor_stmt(ctx *parser.For_stmtContext) interface{} {

// 	v.enterBlock()

// 	// For loop with for-clause
// 	if ctx.For_clause_block() != nil {

// 		v.Visit(ctx.For_clause_block())

// 	}

// 	v.exitBlock()

// 	return nil

// }

// // VisitFor_clause_block evaluates Flo for loop clauses and the block
// func (v *FloVisitor) VisitFor_clause_block(ctx *parser.For_clause_blockContext) interface{} {

// 	v.Visit(ctx.Simple_stmt(0))

// 	for v.Visit(ctx.Expression()).(FloBool) {
// 		v.Visit(ctx.Block())
// 		v.Visit(ctx.Simple_stmt(1))
// 	}

// 	return nil

// }

// // VisitFunc_decl evaluates Flo function definitions
// func (v *FloVisitor) VisitFunc_decl(ctx *parser.Func_declContext) interface{} {

// 	name := ctx.IDENTIFIER()
// 	params := v.Visit(ctx.Parameters())

// 	var args FloList
// 	if params != nil {
// 		args = FloList(params.(FloList))
// 	}

// 	val := FloCallable{
// 		name:      name.GetText(),
// 		min_scope: v.blockScope,
// 		args:      args,
// 		body:      ctx.Block(),
// 	}

// 	v.storeIdentifier(name.GetText(), val)

// 	return nil

// }

// // VisitParameters evaluates Flo callable parameters
// func (v *FloVisitor) VisitParameters(ctx *parser.ParametersContext) interface{} {

// 	vals := make([]FloObject, 0)

// 	for _, val := range ctx.AllIDENTIFIER() {
// 		vals = append(vals, FloString(val.GetText()))
// 	}

// 	return FloList(vals)
// }

// // VisitPrint_stmt executes 'print'
// func (v *FloVisitor) VisitPrint_stmt(ctx *parser.Print_stmtContext) interface{} {
// 	if ctx.Expression() != nil {
// 		x := v.Visit(ctx.Expression())
// 		fmt.Println(x)
// 	}
// 	return nil
// }

// // VisitExpressionGroup is used to evaluate expression groups
// func (v *FloVisitor) VisitExpressionGroup(ctx *parser.ExpressionGroupContext) interface{} {

// 	return ctx.Expression().Accept(v)
// }

// // VisitMulDiv evaluates multiplication and division nodes in the parse tree
// func (v *FloVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserMULTIPLY:
// 		return Multiply(left, right)
// 	case parser.FloParserDIVIDE:
// 		return Divide(left, right)
// 	case parser.FloParserMOD:
// 		return Mod(left, right)
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitAddSub evaluates addition and subtraction nodes in the parse tree
// func (v *FloVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserADDITION:
// 		return Plus(left, right)
// 	case parser.FloParserSUBTRACTION:
// 		return Minus(left, right)
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitBitShift evaluates << and >> nodes in the parse tree
// func (v *FloVisitor) VisitBitShift(ctx *parser.BitShiftContext) interface{} {
// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserLSHIFT:
// 		return LShift(left, right)
// 	case parser.FloParserRSHIFT:
// 		return RShift(left, right)
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitRelational evaluates <, <=, >, >=, ==, !=
// func (v *FloVisitor) VisitRelational(ctx *parser.RelationalContext) interface{} {
// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserLESS:
// 		return LessThan(left, right)
// 	case parser.FloParserLESS_OR_EQUALS:
// 		return LessThanEqual(left, right)
// 	case parser.FloParserGREATER:
// 		return GreaterThan(left, right)
// 	case parser.FloParserGREATER_OR_EQUALS:
// 		return GreaterThanEqual(left, right)
// 	case parser.FloParserEQUALS:
// 		return Equals(left, right)
// 	case parser.FloParserNOT_EQUALS:
// 		return NotEquals(left, right)
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitPower evaluates **
// func (v *FloVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	return Power(left, right)
// }

// // VisitUnary evaluates unary +, -
// func (v *FloVisitor) VisitUnary(ctx *parser.UnaryContext) interface{} {
// 	value := v.Visit(ctx.Expression()).(FloObject)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserADDITION:
// 		return UnaryPlus(value)
// 	case parser.FloParserSUBTRACTION:
// 		return UnaryMinus(value)
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitUnaryIncDec evaluates unary ++, --
// func (v *FloVisitor) VisitUnaryIncDec(ctx *parser.UnaryIncDecContext) interface{} {

// 	// value := v.Visit(ctx.Expression()).(FloObject)
// 	name := ctx.GetStart().GetText()
// 	val := v.readIdentifier(name)

// 	switch ctx.GetOp().GetTokenType() {
// 	case parser.FloParserPLUS_PLUS:
// 		new_val := UnaryPlusPlus(val)
// 		v.storeIdentifier(name, new_val)
// 		return new_val
// 	case parser.FloParserMINUS_MINUS:
// 		new_val := UnaryMinusMinus(val)
// 		v.storeIdentifier(name, new_val)
// 		return new_val
// 	default:
// 		v.doError(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
// 		return nil
// 	}
// }

// // VisitBitwiseAnd evaluates &
// func (v *FloVisitor) VisitBitwiseAnd(ctx *parser.BitwiseAndContext) interface{} {

// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	x := BitwiseAnd(left, right)
// 	return x
// }

// // VisitXOR evaluates ^
// func (v *FloVisitor) VisitXOR(ctx *parser.XORContext) interface{} {

// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	x := BitwiseXOR(left, right)
// 	return x
// }

// // VisitBitwiseOr evaluates |
// func (v *FloVisitor) VisitBitwiseOr(ctx *parser.BitwiseOrContext) interface{} {

// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	x := BitwiseOr(left, right)
// 	return x
// }

// // VisitAnd evaluates logical 'and'
// func (v *FloVisitor) VisitAnd(ctx *parser.AndContext) interface{} {

// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	x := And(left, right)
// 	return x
// }

// // VisitOr evaluates logical 'or'
// func (v *FloVisitor) VisitOr(ctx *parser.OrContext) interface{} {

// 	left, right := v.Visit(ctx.Expression(0)).(FloObject), v.Visit(ctx.Expression(1)).(FloObject)
// 	x := Or(left, right)
// 	return x
// }

// // VisitNot evaluates logical 'not'
// func (v *FloVisitor) VisitNot(ctx *parser.NotContext) interface{} {

// 	value := v.Visit(ctx.Expression()).(FloObject)
// 	x := Not(value)
// 	return x
// }

// // VisitExpressionList evaluates Flo lists
// func (v *FloVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {

// 	list := make([]FloObject, 0)

// 	if ctx.Expression_list() != nil {
// 		vals := v.Visit(ctx.Expression_list()).([]FloObject)
// 		list = append(list, vals...)
// 	}

// 	return FloList(list)

// }

// // VisitExpression_list evaluates Flo expression lists
// func (v *FloVisitor) VisitExpression_list(ctx *parser.Expression_listContext) interface{} {

// 	vals := make([]FloObject, 0)

// 	for _, val := range ctx.AllExpression() {

// 		vals = append(vals, v.Visit(val).(FloObject))

// 	}

// 	return FloList(vals)
// }

// // VisitGetItem evaluates Flo list indicies
// func (v *FloVisitor) VisitGetItem(ctx *parser.GetItemContext) interface{} {

// 	// lval
// 	lval := v.Visit(ctx.Expression(0))
// 	switch lval.(type) {
// 	case FloList:
// 		break
// 	default:
// 		v.doError(listTypeError(lval.(FloObject)))
// 	}
// 	// Index
// 	index := v.Visit(ctx.Expression(1))
// 	switch index.(type) {
// 	case FloInteger:
// 		break
// 	default:
// 		v.doError(listIndexTypeError(index.(FloObject)))
// 	}

// 	list := lval.(FloList)
// 	list_index := index.(FloInteger)

// 	if int(list_index) >= len(list) || int(list_index) < 0 {
// 		v.doError(indexOutOfRangeError())
// 	}

// 	return list[list_index]
// }

// func (v *FloVisitor) callCallable(fn FloCallable, args FloList) FloObject {

// 	prev_func_scope := v.funcScope

// 	v.push(fn)

// 	v.enterBlock()

// 	v.funcScope = v.blockScope

// 	// Check arity
// 	if len(args) != len(fn.args) {
// 		v.doError(callableArgsTypeError(fn, len(args)))
// 	}

// 	// Initialise callable with parameters
// 	for i, arg := range args {
// 		v.storeIdentifier(string(fn.args[i].(FloString)), arg)
// 	}

// 	// Call function
// 	fn.body.Accept(v)

// 	// Pop return value off call stack
// 	x := v.pop()

// 	// Pop function off call stack
// 	v.pop()

// 	v.exitBlock()

// 	v.funcScope = prev_func_scope

// 	if x == nil {
// 		return nil
// 	}
// 	// fmt.Println(v.Stack)
// 	return x.(FloObject)
// }

// // VisitCallExpression evaluates Flo function calls
// func (v *FloVisitor) VisitCallExpression(ctx *parser.CallExpressionContext) interface{} {

// 	// if ctx.Expression() == nil {
// 	// 	v.doError("n")
// 	// }

// 	x := v.Visit(ctx.Expression())

// 	switch x.(type) {
// 	case FloCallable:
// 		break
// 	default:
// 		v.doError(notCallableError(x.(FloObject)))
// 	}

// 	fn := x.(FloCallable)
// 	y := v.Visit(ctx.Expression_list())
// 	var args FloList
// 	switch y.(type) {
// 	case FloList:
// 		args = y.(FloList)
// 		break
// 	case nil:
// 		break
// 	default:
// 		v.doError(notCallableError(y.(FloObject)))
// 	}

// 	return v.callCallable(fn, args)

// }

// // VisitReturn_stmt evaluates Flo function calls
// func (v *FloVisitor) VisitReturn_stmt(ctx *parser.Return_stmtContext) interface{} {

// 	if ctx.Expression() == nil {
// 		return nil
// 	}

// 	x := v.Visit(ctx.Expression())
// 	v.push(x.(FloObject))
// 	return nil

// }

// // VisitInteger evaluates Flo integers
// func (v *FloVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {

// 	i, err := strconv.Atoi(ctx.GetText())
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return FloInteger(i)

// }

// // VisitFloat evaluates floats in Flo
// func (v *FloVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {

// 	f, err := strconv.ParseFloat(ctx.GetText(), 64)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return FloFloat(f)
// }

// // VisitString evaluates Flo strings
// func (v *FloVisitor) VisitString(ctx *parser.StringContext) interface{} {

// 	value := ctx.GetText()
// 	value = value[1 : len(value)-1]
// 	return FloString(value)

// }

// // VisitBoolean evaluates Flo booleans
// func (v *FloVisitor) VisitBoolean(ctx *parser.BooleanContext) interface{} {

// 	x := ctx.GetText()
// 	if x == "true" {
// 		return FloBool(true)
// 	}
// 	return FloBool(false)

// }

// // VisitReadIdentifier evaluates Flo identifiers
// func (v *FloVisitor) VisitReadIdentifier(ctx *parser.ReadIdentifierContext) interface{} {

// 	identifier := ctx.GetText()
// 	val := v.readIdentifier(identifier)
// 	if val != nil {
// 		return val
// 	}
// 	return nil

// }

// // VisitBlock evaluates blocks
// func (v *FloVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {

// 	v.enterBlock()
// 	if ctx.Multi_stmts() != nil {
// 		v.Visit(ctx.Multi_stmts())
// 	}
// 	v.exitBlock()

// 	return nil
// }

// // Visit a parse tree produced by FloParser#eos.
// func (v *FloVisitor) VisitEos(ctx *parser.EosContext) interface{} {
// 	return nil
// }

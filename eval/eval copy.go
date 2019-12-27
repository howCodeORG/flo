package eval

// import (
// 	"flo/parser"
// 	"fmt"
// 	"strconv"
// )

// // FloListener holds the ANTLR stack
// type FloListener struct {
// 	*parser.BaseFloListener

// 	stack []FloObject

// 	scope       int
// 	environment []map[string]FloObject
// 	Error       bool
// }

// func (l *FloListener) cleanup() {
// 	l.scope = 0
// }

// // Init creates the environment for Flo program state
// func (l *FloListener) Init() {
// 	l.environment = make([]map[string]FloObject, 1)
// 	l.environment[l.scope] = make(map[string]FloObject, 1)
// }

// func (l *FloListener) push(i FloObject) {
// 	l.stack = append(l.stack, i)
// }

// func (l *FloListener) pop() FloObject {

// 	if len(l.stack) < 1 {
// 		panic("stack is empty unable to pop")
// 	}

// 	// Get the last value from the stack.
// 	result := l.stack[len(l.stack)-1]

// 	// Remove the last element from the stack.
// 	l.stack = l.stack[:len(l.stack)-1]

// 	return result
// }

// // Out returns the output of an expression
// func (l *FloListener) Out() FloObject {

// 	if len(l.stack) >= 1 {
// 		return l.pop()
// 	}
// 	return nil
// }

// func (l *FloListener) doError(err string) {
// 	l.cleanup()
// 	l.stack = make([]FloObject, 0)
// 	panic(err)
// }

// // ExitMulDiv evaluates multiplication and division nodes in the parse tree
// func (l *FloListener) ExitMulDiv(c *parser.MulDivContext) {
// 	right, left := l.pop(), l.pop()

// 	switch c.GetOp().GetTokenType() {
// 	case parser.FloParserMULTIPLY:
// 		l.push(Multiply(left, right))
// 	case parser.FloParserDIVIDE:
// 		l.push(Divide(left, right))
// 	case parser.FloParserMOD:
// 		l.push(Mod(left, right))
// 	default:
// 		l.doError(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
// 	}
// }

// // ExitAddSub evaluates addition and subtraction nodes in the parse tree
// func (l *FloListener) ExitAddSub(c *parser.AddSubContext) {
// 	right, left := l.pop(), l.pop()

// 	switch c.GetOp().GetTokenType() {
// 	case parser.FloParserADDITION:
// 		l.push(Plus(left, right))
// 	case parser.FloParserSUBTRACTION:
// 		l.push(Minus(left, right))
// 	default:
// 		l.doError(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
// 	}
// }

// // ExitBitShift evaluates << and >>
// func (l *FloListener) ExitBitShift(c *parser.BitShiftContext) {
// 	right, left := l.pop(), l.pop()

// 	switch c.GetOp().GetTokenType() {
// 	case parser.FloParserLSHIFT:
// 		l.push(LShift(left, right))
// 	case parser.FloParserRSHIFT:
// 		l.push(RShift(left, right))
// 	default:
// 		l.doError(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
// 	}
// }

// // ExitRelational evaluates <, <=, >, >=, ==, !=
// func (l *FloListener) ExitRelational(c *parser.RelationalContext) {
// 	right, left := l.pop(), l.pop()

// 	switch c.GetOp().GetTokenType() {
// 	case parser.FloParserLESS:
// 		l.push(LessThan(left, right))
// 	case parser.FloParserLESS_OR_EQUALS:
// 		l.push(LessThanEqual(left, right))
// 	case parser.FloParserGREATER:
// 		l.push(GreaterThan(left, right))
// 	case parser.FloParserGREATER_OR_EQUALS:
// 		l.push(GreaterThanEqual(left, right))
// 	case parser.FloParserEQUALS:
// 		l.push(Equals(left, right))
// 	case parser.FloParserNOT_EQUALS:
// 		l.push(NotEquals(left, right))
// 	default:
// 		l.doError(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
// 	}
// }

// // ExitPower evaluates **
// func (l *FloListener) ExitPower(c *parser.PowerContext) {

// 	right, left := l.pop(), l.pop()

// 	l.push(Power(left, right))
// }

// // ExitUnary evaluates unary addition and subtraction
// func (l *FloListener) ExitUnary(c *parser.UnaryContext) {
// 	x := l.pop()
// 	switch c.GetOp().GetTokenType() {
// 	case parser.FloParserADDITION:
// 		l.push(UnaryPlus(x))
// 	case parser.FloParserSUBTRACTION:
// 		l.push(UnaryMinus(x))
// 	default:
// 		l.doError(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
// 	}
// }

// // ExitBitwiseAnd evaluates Bitwise and's
// func (l *FloListener) ExitBitwiseAnd(c *parser.BitwiseAndContext) {

// 	right, left := l.pop(), l.pop()

// 	x := BitwiseAnd(left, right)
// 	l.push(FloObject(x))

// }

// // ExitXOR evaluates Bitwise XOR's
// func (l *FloListener) ExitXOR(c *parser.XORContext) {

// 	right, left := l.pop(), l.pop()

// 	x := BitwiseXOR(left, right)
// 	l.push(FloObject(x))

// }

// // ExitBitwiseOr evaluates Bitwise Or's
// func (l *FloListener) ExitBitwiseOr(c *parser.BitwiseOrContext) {

// 	right, left := l.pop(), l.pop()

// 	x := BitwiseOr(left, right)
// 	l.push(FloObject(x))

// }

// // ExitAnd evaluates logical and's
// func (l *FloListener) ExitAnd(c *parser.AndContext) {

// 	right, left := l.pop(), l.pop()

// 	x := And(left, right)
// 	l.push(FloBool(x))

// }

// // ExitOr evaluates logical or's
// func (l *FloListener) ExitOr(c *parser.OrContext) {

// 	right, left := l.pop(), l.pop()

// 	x := Or(left, right)

// 	l.push(FloBool(x))

// }

// // ExitNot evaluates logical not's
// func (l *FloListener) ExitNot(c *parser.NotContext) {

// 	val := l.pop()

// 	x := Not(val)

// 	l.push(FloBool(x))

// }

// // ExitInteger evaluates integers
// func (l *FloListener) ExitInteger(c *parser.IntegerContext) {
// 	i, err := strconv.Atoi(c.GetText())
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	l.push(FloInteger(i))
// }

// // ExitFloat evaluates floats in Flo
// func (l *FloListener) ExitFloat(c *parser.FloatContext) {
// 	i, err := strconv.ParseFloat(c.GetText(), 64)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	l.push(FloFloat(i))
// }

// // ExitReadIdentifier evaluates Flo identifiers
// func (l *FloListener) ExitReadIdentifier(c *parser.ReadIdentifierContext) {

// 	identifier := c.GetText()
// 	// Lookup identifier value
// 	var found bool
// 	for i := l.scope; i >= 0; i-- {
// 		if value, ok := l.environment[i][identifier]; ok {
// 			l.push(value)
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		l.doError("Undefined variable")
// 	}

// }

// // ExitStoreIdentifier evaluates Flo identifiers
// func (l *FloListener) ExitStoreIdentifier(c *parser.StoreIdentifierContext) {

// 	identifier := c.GetStart().GetText()

// 	// Lookup identifier value
// 	l.environment[l.scope][identifier] = l.pop()

// }

// // EnterBlock increases the current scope
// func (l *FloListener) EnterBlock(c *parser.BlockContext) {
// 	// Increase scope
// 	l.scope++
// 	if l.scope == len(l.environment) {
// 		l.environment = append(l.environment, make(map[string]FloObject, 1))
// 	}
// }

// // ExitBlock increases the current scope
// func (l *FloListener) ExitBlock(c *parser.BlockContext) {
// 	// Decrease scope
// 	l.scope--
// 	l.environment = l.environment[:len(l.environment)-1]
// }

// // ExitString evaluates strings in Flo
// func (l *FloListener) ExitString(c *parser.StringContext) {

// 	x := c.GetText()
// 	// Remove quotes
// 	x = x[1 : len(x)-1]
// 	l.push(FloString(x))

// }

// // ExitBoolean evaluates bools in Flo
// func (l *FloListener) ExitBoolean(c *parser.BooleanContext) {

// 	x := c.GetText()
// 	if x == "true" {
// 		l.push(FloBool(true))
// 	} else {
// 		l.push(FloBool(false))
// 	}

// }

// // ExitPrintStatement prints stuff to the console
// func (l *FloListener) ExitPrintStatement(c *parser.PrintStatementContext) {

// 	out := l.Out()
// 	if out != nil {
// 		fmt.Println(out)
// 	}

// }

// // ExitIfStatement evaluates if statements
// func (l *FloListener) ExitIfStatement(c *parser.IfStatementContext) {

// 	x := c.GetText()

// 	fmt.Println(x)

// }

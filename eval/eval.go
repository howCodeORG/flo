package eval

import (
	"bufio"
	"flo/ast"
	"flo/object"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// ast.Node
// ast.List
// ast.Statement
// ast.UnaryOp
// ast.BinaryOp
// ast.Function

type env map[string]object.Object

var environment []env
var callStack []object.Object
var scope int
var returnFlag bool
var breakFlag bool
var Error string

func callStackPush(o object.Object) {
	callStack = append(callStack, o)
	// fmt.Println("push call stack: ", callStack)
}

func callStackPop() object.Object {

	if len(callStack) > 0 {
		value := callStack[len(callStack)-1]
		callStack = callStack[:len(callStack)-1]
		return value
	}
	return nil

}

func newScope() {
	environment = append(environment, make(env))
	scope++
}

func removeScope() {
	environment = environment[:len(environment)-1]
	scope--
}

func incScope() {
	scope++
}

func decScope() {
	scope--
}

func setError(err string) {
	if Error == "" {
		Error = err
	}
}

func foundError() bool {
	if Error != "" {
		return true
	}
	return false
}

func addBuiltinType() {
	environment[0]["type"] = &object.Function{
		Name:    "type",
		Builtin: true,
		Parameters: []ast.Noder{
			&ast.Node{Type: "IDENT", Value: "var"},
		},
		Code: func(params ...ast.Noder) ast.Noder {

			// Only expecting 1 parameter for the type() function

			pt := Eval(params[0])
			//fmt.Println(pt.Type(), pt.String())
			parameter_type := "void"
			if pt != nil {
				parameter_type = pt.Type()
			}

			return &ast.Node{
				Type:  "STRING",
				Value: parameter_type,
			}
		},
	}
}

func addBuiltinOpen() {
	environment[0]["open"] = &object.Function{
		Name:    "open",
		Builtin: true,
		Parameters: []ast.Noder{
			&ast.Node{Type: "IDENT", Value: "filename"},
			// &ast.Node{Type: "IDENT", Value: "mode"},
		},
		Code: func(params ...ast.Noder) ast.Noder {

			// Only expecting 2 parameters for the open() function
			// Second parameter disabled for now

			filename := Eval(params[0])
			if filename == nil {
				return &ast.Node{
					Type:  "INT",
					Value: "-1",
				}
			}
			// mode := Eval(params[1])
			// if mode == nil {
			// 	return &ast.Node{
			// 		Type:  "INT",
			// 		Value: "-1",
			// 	}
			// }

			file, err := os.Open(filename.String())
			if err != nil {
				return &ast.Node{
					Type:  "INT",
					Value: "-1",
				}
			}
			b, err := ioutil.ReadAll(file)
			file_data := string(b)

			defer file.Close()

			return &ast.Node{
				Type:  "STRING",
				Value: file_data,
			}
		},
	}
}

func addBuiltinInput() {
	environment[0]["input"] = &object.Function{
		Name:    "input",
		Builtin: true,
		Parameters: []ast.Noder{
			&ast.Node{Type: "IDENT", Value: "text"},
		},
		Code: func(params ...ast.Noder) ast.Noder {

			// Only expecting 1 parameter for the input() function

			text := Eval(params[0])
			fmt.Print(text)
			reader := bufio.NewReader(os.Stdin)
			input_text, err := reader.ReadString('\n')

			if err != nil {
				return &ast.Node{
					Type:  "INT",
					Value: "-1",
				}
			}
			return &ast.Node{
				Type:  "STRING",
				Value: strings.TrimSuffix(input_text, "\n"),
			}
		},
	}
}

func addBuiltinInt() {
	environment[0]["int"] = &object.Function{
		Name:    "int",
		Builtin: true,
		Parameters: []ast.Noder{
			&ast.Node{Type: "IDENT", Value: "value"},
		},
		Code: func(params ...ast.Noder) ast.Noder {
			// Only expecting 1 parameter for the input() function
			val := Eval(params[0])

			if val.Type() == object.INTEGER_OBJECT {
				return &ast.Node{
					Type:  "INT",
					Value: val.String(),
				}
			}

			if val.Type() == object.STRING_OBJECT {
				_, err := strconv.Atoi(strings.Split(val.String(), ".")[0])
				if err != nil {
					return &ast.Node{
						Type:  "NIL",
						Value: "nil",
					}
				}
				return &ast.Node{
					Type:  "INT",
					Value: strings.Split(val.String(), ".")[0],
				}
			}

			if val.Type() == object.FLOAT_OBJECT {
				_, err := strconv.Atoi(strings.Split(val.String(), ".")[0])
				if err != nil {
					fmt.Println(err)
					return &ast.Node{
						Type:  "NIL",
						Value: "nil",
					}
				}
				return &ast.Node{
					Type:  "INT",
					Value: strings.Split(val.String(), ".")[0],
				}
			}
			return &ast.Node{
				Type:  "NIL",
				Value: "nil",
			}
		},
	}
}

func addBuiltins() {
	// Built-in type() function
	// Name: type
	// Parameters: 1
	// Returns: The type of the object passed to it
	addBuiltinType()

	// Built-in open() function
	// Name: open
	// Parameters: 2
	// Returns: A string of the data from a file
	addBuiltinOpen()

	// Built-in input() function
	// Name: input
	// Parameters: 1
	// Returns: A string from stdin
	addBuiltinInput()

	// Built-in int() function
	// Name: int
	// Parameters: 1
	// Returns: An int object
	addBuiltinInt()
}

func Init() {
	environment = make([]env, 1)
	environment[0] = make(env)

	addBuiltins()

	scope = 0
	returnFlag = false
	Error = ""
}

func Eval(node ast.Noder) object.Object {

	if Error != "" || node == nil {
		return nil
	}

	if returnFlag == false {
		switch n := node.(type) {
		case *ast.Block:
			return evalBlock(n)
		case *ast.Node:
			return evalNode(n)
		case *ast.List:
			return evalList(n)
		case *ast.Statement:
			return evalStatement(n)
		case *ast.UnaryOp:
			return evalUnaryOp(n)
		case *ast.BinaryOp:
			return evalBinaryOp(n)
		case *ast.IfStatement:
			return evalIfStatement(n)
		case *ast.ForStatement:
			return evalForStatement(n)
		case *ast.Function:
			return evalFunctionDeclaration(n)
		}
		return nil
	}
	//fmt.Println("node:", node, "callstack:", callStack)
	// return callStackPop()
	return nil
}

func evalBlock(node *ast.Block) object.Object {

	if Error != "" {
		return nil
	}

	newScope()
	output := Eval(node.Value)
	removeScope()
	return output
}

func evalStatement(node *ast.Statement) object.Object {

	if Error != "" {
		return nil
	}
	var output object.Object
	if node.Value != nil {
		output = Eval(node.Value)
		if output != nil {
			// fmt.Println(output)
		}
	}
	if node.Next != nil {
		output = Eval(node.Next)
		if output != nil {
			// fmt.Println(output)
		}
	}

	return output
}

func evalNode(node *ast.Node) object.Object {

	if Error != "" {
		return nil
	}

	switch node.Type {
	case "INT":
		value, err := strconv.Atoi(node.Value)
		if err != nil {
			return nil
		}
		return &object.Integer{Value: value}
	case "FLOAT":
		value, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return nil
		}
		return &object.Float{Value: value}
	case "STRING":
		return &object.String{Value: node.Value}
	case "BOOL":
		value, err := strconv.ParseBool(node.Value)
		if err != nil {
			return nil
		}
		return &object.Bool{Value: value}
	case "IDENT":
		value, _ := getIdentifier(node.Value)
		return value
	default:
		return &object.Nil{}
	}
}

func getIdentifier(name string) (object.Object, bool) {

	currentScope := len(environment) - 1
	for i := currentScope; i >= 0; i-- {
		if val, ok := environment[i][name]; ok {
			return val, true
		}
	}
	Error = "variable is not defined"
	return nil, false

}

func evalFunctionDeclaration(node *ast.Function) object.Object {

	if Error != "" {
		return nil
	}

	name := node.Name
	parameters := node.Parameters
	code := node.Body

	func_scope := scope

	// Anonymous functions are accessible from the global scope
	// because it makes the implementation easier and
	// nobody can access them directly anyway due
	// to '@' being illegal in function names.
	if name[0] == '@' {
		func_scope = 0
	}

	environment[func_scope][name] = &object.Function{Builtin: false, Name: name, Parameters: parameters.Values, Code: func(params ...ast.Noder) ast.Noder { return Eval(code) }}

	return environment[func_scope][name]
}

func evalList(node *ast.List) object.Object {

	if Error != "" {
		return nil
	}
	l := &object.List{}
	for _, item := range node.Values {
		l.Values = append(l.Values, Eval(item))
	}
	return l
}

func evalUnaryOp(node *ast.UnaryOp) object.Object {

	if Error != "" {
		return nil
	}

	switch node.Operator {
	case "+":
		return unaryAddOp(Eval(node.Value))
	case "-":
		return unarySubOp(Eval(node.Value))
	case "not":
		return unaryNotOp(Eval(node.Value))
	case "return":
		return unaryReturnOp(Eval(node.Value))
	case "break":
		return unaryBreakOp()
	case "print":
		return unaryPrintOp(Eval(node.Value))
	default:
		return nil
	}

}

func unaryAddOp(o object.Object) object.Object {
	// Check types are compatible
	// + int
	if o.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: (o.(*object.Integer)).Value}
	}
	// + float
	if o.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: (o.(*object.Float)).Value}
	}
	// // not bool
	// if o.Type() == object.BOOLEAN_OBJECT {
	// 	return &object.Bool{Value: !(o.(*object.Bool)).Value}
	// }
	return nil
}

func unarySubOp(o object.Object) object.Object {
	// Check types are compatible
	// - int
	if o.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: (o.(*object.Integer)).Value * -1}
	}
	// - float
	if o.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: (o.(*object.Float)).Value * -1}
	}
	return nil
}

func unaryNotOp(o object.Object) object.Object {
	// Check types are compatible
	// not bool
	if o.Type() == object.BOOLEAN_OBJECT {
		return &object.Bool{Value: !(o.(*object.Bool)).Value}
	}
	return nil
}

func unaryReturnOp(o object.Object) object.Object {

	// fmt.Println("o:", o, "before:", callStack)

	// Set returning flag to prevent any more code being executed
	returnFlag = true
	// Replace top of call stack with expression
	// callStack[len(callStack)-1] = o
	callStackPop()
	callStackPush(o)
	// fmt.Println("after:", callStack)

	// Pop top of call stack
	return o

}

func unaryBreakOp() object.Object {

	// Set break flag to exit loops
	breakFlag = true

	return nil

}

func unaryPrintOp(o object.Object) object.Object {

	if o != nil {
		fmt.Println(o)
	}
	return nil

}

func evalBinaryOp(node *ast.BinaryOp) object.Object {

	if Error != "" {
		return nil
	}

	switch node.Operator {
	case "+":
		return binaryAddOp(Eval(node.Left), Eval(node.Right))
	case "-":
		return binarySubOp(Eval(node.Left), Eval(node.Right))
	case "*":
		return binaryMulOp(Eval(node.Left), Eval(node.Right))
	case "/":
		return binaryDivOp(Eval(node.Left), Eval(node.Right))
	case "**":
		return binaryPowOp(Eval(node.Left), Eval(node.Right))
	case "%":
		return binaryModOp(Eval(node.Left), Eval(node.Right))
	case "=":
		return binaryAssignOp(node.Left, node.Right)
	case "()":
		return binaryFunctionOp(node.Left, node.Right)
	case "[]":
		return binaryListOp(node.Left, node.Right)
	case "==":
		return binaryEqeqOp(node.Left, node.Right)
	case "!=":
		return binaryNteqOp(node.Left, node.Right)
	case "<=":
		return binaryLteqOp(node.Left, node.Right)
	case ">=":
		return binaryGteqOp(node.Left, node.Right)
	case "<":
		return binaryLtOp(node.Left, node.Right)
	case ">":
		return binaryGtOp(node.Left, node.Right)
	default:
		return nil
	}
}

func binaryAddOp(left object.Object, right object.Object) object.Object {

	// Check types are compatible
	// int + int = int
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: left.(*object.Integer).Value + right.(*object.Integer).Value}
	}
	// int + float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: float64(left.(*object.Integer).Value) + right.(*object.Float).Value}
	}
	// float + int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value + float64(right.(*object.Integer).Value)}
	}
	// float + float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value + right.(*object.Float).Value}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for +: '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binarySubOp(left object.Object, right object.Object) object.Object {
	// Check types are compatible
	// int - int = int
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: left.(*object.Integer).Value - right.(*object.Integer).Value}
	}
	// int - float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: float64(left.(*object.Integer).Value) - right.(*object.Float).Value}
	}
	// float - int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value - float64(right.(*object.Integer).Value)}
	}
	// float - float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value - right.(*object.Float).Value}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for -: '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binaryMulOp(left object.Object, right object.Object) object.Object {
	// Check types are compatible
	// int * int = int
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: left.(*object.Integer).Value * right.(*object.Integer).Value}
	}
	// int * float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: float64(left.(*object.Integer).Value) * right.(*object.Float).Value}
	}
	// float * int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value * float64(right.(*object.Integer).Value)}
	}
	// float * float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value * right.(*object.Float).Value}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for *: '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binaryDivOp(left object.Object, right object.Object) object.Object {
	// Check types are compatible
	// int / int = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: left.(*object.Integer).Value / right.(*object.Integer).Value}
	}
	// int / float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: float64(left.(*object.Integer).Value) / right.(*object.Float).Value}
	}
	// float / int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value / float64(right.(*object.Integer).Value)}
	}
	// float / float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: left.(*object.Float).Value / right.(*object.Float).Value}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for /: '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binaryPowOp(left object.Object, right object.Object) object.Object {
	// Check types are compatible
	// int ** int = int
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: int(math.Pow(float64(left.(*object.Integer).Value), float64(right.(*object.Integer).Value)))}
	}
	// int ** float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: math.Pow(float64(left.(*object.Integer).Value), float64(right.(*object.Float).Value))}
	}
	// float ** int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: math.Pow(left.(*object.Float).Value, float64(right.(*object.Integer).Value))}
	}
	// float ** float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: math.Pow(left.(*object.Float).Value, float64(right.(*object.Float).Value))}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for **: '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binaryModOp(left object.Object, right object.Object) object.Object {
	// Check types are compatible
	// int % int = int
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Integer{Value: int(math.Mod(float64(left.(*object.Integer).Value), float64(right.(*object.Integer).Value)))}
	}
	// int % float = float
	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: math.Mod(float64(left.(*object.Integer).Value), float64(right.(*object.Float).Value))}
	}
	// float % int = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Float{Value: math.Mod(left.(*object.Float).Value, float64(right.(*object.Integer).Value))}
	}
	// float % float = float
	if left.Type() == object.FLOAT_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Float{Value: math.Mod(left.(*object.Float).Value, float64(right.(*object.Float).Value))}
	}
	setError(fmt.Sprintf("TypeError: unsupported type(s) for %% : '%s' and '%s'", left.Type(), right.Type()))
	return nil
}

func binaryAssignOp(left ast.Noder, right ast.Noder) object.Object {

	varNames := (left.(*ast.List)).Values
	varValues := (right.(*ast.List)).Values

	for i, _ := range varNames {
		name := varNames[i].String()
		result := Eval(varValues[i])
		environment[scope][name] = result
	}

	return nil

}

func binaryFunctionOp(name ast.Noder, params ast.Noder) object.Object {

	function := Eval(name)

	if function.Type() != object.FUNCTION_OBJECT {
		setError(fmt.Sprintf("TypeError: can't use () operator on type: %s", function.Type()))
		return nil
	}

	// fmt.Println("env", environment)

	// Get function name first
	functionName := (function.(*object.Function)).Name

	// Get params
	functionParams := []ast.Noder{}
	if params != nil {
		functionParams = (params.(*ast.List)).Values
	}

	// Get declared function from environment
	declaredFunction, ok := getIdentifier(functionName)
	if ok == true {
		callStackPush(declaredFunction)
		declaredParams := (declaredFunction.(*object.Function)).Parameters

		newScope()

		if len(declaredParams) != len(functionParams) {
			setError(fmt.Sprintf("function needs %d parameters, got %d", len(declaredParams), len(functionParams)))
			removeScope()
			return nil
		}

		for i, _ := range declaredParams {
			// fmt.Println(declaredParams[i], functionParams[i])
			if (declaredFunction.(*object.Function)).Builtin == false {
				environment[scope][declaredParams[i].String()] = Eval(functionParams[i])
			}
		}

		decScope()

		// Execute function
		// fmt.Println("callstack before:", callStack)
		toCall := callStackPop()
		var fn object.Object
		if toCall.Type() == object.FUNCTION_OBJECT {
			fn = Eval((toCall.(*object.Function)).Code(functionParams...))
		} else {
			fn = toCall
		}
		// fmt.Println("callstack after:", callStack)
		// Update returning flag to false
		returnFlag = false

		incScope()

		removeScope()
		if fn == nil {
			if len(callStack) > 0 {
				return callStack[len(callStack)-1]
			}
			return &object.Nil{}
		}
		return fn
	}
	setError(fmt.Sprintf("undeclared function"))
	return nil

}

func binaryListOp(name ast.Noder, param ast.Noder) object.Object {

	list_name := name.String()
	list_param := Eval(param)
	list_var, ok := getIdentifier(list_name)
	if ok == true {
		index, err := strconv.Atoi(list_param.String())
		if err != nil {
			setError(fmt.Sprintf("something went wrong"))
			return nil
		}

		list_vals := (list_var.(*object.List)).Values

		if index < 0 || index > len(list_vals)-1 {
			setError(fmt.Sprintf("IndexError: out of range"))
			return nil
		}
		return list_vals[index]

	}

	return nil

}

func binaryEqeqOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value == right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value != float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) == right.(*object.Float).Value}
	} else if left.Type() == object.STRING_OBJECT && right.Type() == object.STRING_OBJECT {
		return &object.Bool{Value: left.(*object.String).Value == right.(*object.String).Value}
	} else if left.Type() == object.BOOLEAN_OBJECT && right.Type() == object.BOOLEAN_OBJECT {
		return &object.Bool{Value: left.(*object.Bool).Value == right.(*object.Bool).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for ==: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func binaryNteqOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value != right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value != float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) != right.(*object.Float).Value}
	} else if left.Type() == object.BOOLEAN_OBJECT && right.Type() == object.BOOLEAN_OBJECT {
		return &object.Bool{Value: left.(*object.Bool).Value != right.(*object.Bool).Value}
	} else if left.Type() == object.STRING_OBJECT && right.Type() == object.STRING_OBJECT {
		return &object.Bool{Value: left.(*object.String).Value != right.(*object.String).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for !=: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func binaryLteqOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value <= right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value <= float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) <= right.(*object.Float).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for <=: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func binaryGteqOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value >= right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value >= float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) >= right.(*object.Float).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for >=: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func binaryLtOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value < right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value < float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) < right.(*object.Float).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for <: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func binaryGtOp(l ast.Noder, r ast.Noder) object.Object {

	left := Eval(l)
	right := Eval(r)

	if left.Type() == object.INTEGER_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Integer).Value > right.(*object.Integer).Value}
	} else if left.Type() == object.FLOAT_OBJECT && right.Type() == object.INTEGER_OBJECT {
		return &object.Bool{Value: left.(*object.Float).Value > float64(right.(*object.Integer).Value)}
	} else if left.Type() == object.INTEGER_OBJECT && right.Type() == object.FLOAT_OBJECT {
		return &object.Bool{Value: float64(left.(*object.Integer).Value) > right.(*object.Float).Value}
	} else {
		setError(fmt.Sprintf("TypeError: unsupported type(s) for >: '%s' and '%s'", left.Type(), right.Type()))
		return nil
	}

}

func evalIfStatement(node *ast.IfStatement) object.Object {

	if Error != "" {
		return nil
	}

	obj := Eval(node.Condition)

	if obj == nil {
		return nil
	}

	if obj.Type() == object.BOOLEAN_OBJECT {

		if (obj.(*object.Bool)).Value == true {
			return Eval(node.If)
		} else {

			return Eval(node.Else)
		}

	} else {
		setError(fmt.Sprintf("if condition must evaluate to bool"))
		return nil
	}
}

func evalForStatement(node *ast.ForStatement) object.Object {

	if foundError() {
		return nil
	}

	// Normally scope is increased when we enter a block and decreased when we leave a block,
	// delimited by '{' and '}'. If-statements and for-loops require their condition variables
	// be readable and modifiable from within their blocks.
	// Ex. if x == 1 { x = x + 1; print x }
	//               ^
	// Scope normally starts here, meaning the 'x' inside the braces is a different 'x' to the one
	// being tested.
	// We need to move the scope to here:
	//     if x == 1 { x = x + 1; print x}
	//        ^
	// We do that by incrementing the 'scopeLevel' variable here, and decrementing it once we exit the
	// if-statement or for-loop.

	// Determine what kind of loop we're evaluating first
	if node.Condition == nil {
		// This loop corresponds to parser rule 3 aka the empty for loop, 'for {}'
		// Semantically this loop is equivalent to 'for true {}'
		for true {
			Eval(node.Body)
			if returnFlag == true || breakFlag == true {
				breakFlag = false
				break
			}
		}
	} else {
		// This loop could be rule 1 or rule 2
		// rule 1: for <condition> <block>
		// rule 2: for <init> ';' <condition> ';' <post> <block>

		// Check if init is nil
		if node.Init == nil {

			// If this is true, we're using rule 1, otherwise rule 2
			// Check condition is true
			newScope()
			for Eval(node.Condition).(*object.Bool).Value == true {
				decScope()
				Eval(node.Body)
				incScope()
				if returnFlag == true || breakFlag == true {
					breakFlag = false
					break
				}
			}
			removeScope()

		} else {

			// Rule 2
			// Evaluate Init before the loop
			newScope()
			Eval(node.Init)
			decScope()
			// Evaluate Condition before loop
			// Evaluate Post after each loop iteration

			incScope()
			for Eval(node.Condition).(*object.Bool).Value == true {
				// Evaluate body of loop
				decScope()
				Eval(node.Body)

				// Now evaluate Post
				incScope()
				if returnFlag == true || breakFlag == true {
					breakFlag = false
					break
				}
				Eval(node.Post)

			}
			removeScope()
		}

	}

	return nil

}

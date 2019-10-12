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
		Name: "type",
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
		Name: "open",
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
		Name: "input",
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

func addBuiltinInteger() {
	environment[0]["integer"] = &object.Function{
		Name: "integer",
		Parameters: []ast.Noder{
			&ast.Node{Type: "IDENT", Value: "value"},
		},
		Code: func(params ...ast.Noder) ast.Noder {

			// Only expecting 1 parameter for the input() function

			val := Eval(params[0])

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

	// Built-in integer() function
	// Name: integer
	// Parameters: 1
	// Returns: An int object
	addBuiltinInteger()
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
		case *ast.Statement:
			return evalStatement(n)
		case *ast.UnaryOp:
			return evalUnaryOp(n)
		case *ast.BinaryOp:
			return evalBinaryOp(n)
		case *ast.TrinaryOp:
			return evalTrinaryOp(n)
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

	environment[scope][name] = &object.Function{Name: name, Parameters: parameters.Values, Code: func(params ...ast.Noder) ast.Noder { return Eval(code) }}

	return nil
}

func evalList(node *ast.List) []ast.Noder {

	if Error != "" {
		return nil
	}

	return node.Values
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

	// Get function name first
	functionName := name.String()
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
			environment[scope][declaredParams[i].String()] = Eval(functionParams[i])
		}

		decScope()

		// Execute function
		fn := Eval((callStackPop().(*object.Function)).Code(functionParams...))

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

func evalTrinaryOp(node *ast.TrinaryOp) object.Object {

	if Error != "" {
		return nil
	}

	switch node.Type {
	case "if":
		return trinaryIfOp(Eval(node.Condition), node.If, node.Else)
	default:
		return nil
	}
}

func trinaryIfOp(obj object.Object, if_part ast.Noder, else_part ast.Noder) object.Object {

	if obj == nil {
		return nil
	}

	if obj.Type() == object.BOOLEAN_OBJECT {

		if (obj.(*object.Bool)).Value == true {
			return Eval(if_part)
		} else {

			return Eval(else_part)
		}

	} else {
		setError(fmt.Sprintf("if condition must evaluate to bool"))
		return nil
	}

}

package vm

import (
	"encoding/binary"
	"fmt"
)

type Stack struct {
	data []FloObject
}

// Push pushes a new value onto the stack
func (s *Stack) Push(x FloObject) {
	s.data = append(s.data, x)
}

// Pop removes the top value from the stack
func (s *Stack) Pop() FloObject {
	if len(s.data) > 0 {
		value := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return value
	}
	return nil
}

// Frame is a function call frame
type Frame struct {
	name       string
	dataStack  Stack
	blockStack []map[string]FloObject
}

// VM is the Flo virtual machine
type VM struct {
	callStack []Frame
}

// Virtual machine instructions
const (
	LOAD_CONST = iota
	LOAD_NAME
	STORE_NAME
	EXTENDED_ARG
	PUSH_BLOCK
	POP_BLOCK
	BUILD_LIST
	GET_ITEM
	SETUP_FUNCTION
	SETUP_PARAMS
	SETUP_BODY
	MAKE_FUNCTION
	CALL_FUNCTION
	RETURN
	BINARY_ADD
	BINARY_SUB
	BINARY_DIV
	BINARY_MUL
	BINARY_POW
	BINARY_MOD
	BINARY_LSHIFT
	BINARY_RSHIFT
	BINARY_AND
	BINARY_XOR
	BINARY_OR
	UNARY_NOT
	UNARY_ADD
	UNARY_SUB
	LOGICAL_OR
	LOGICAL_AND
	COMPARE
	POP_JUMP_IF_FALSE
	POP_JUMP_IF_TRUE
	JUMP_IF_FALSE
	JUMP_BACK
	JUMP
	PRINT
	NOP
)

func (vm *VM) tos() *Frame {
	if len(vm.callStack) > 0 {
		return &vm.callStack[len(vm.callStack)-1]
	}
	return nil
}

func (vm *VM) funcCall(name string) {
	f := Frame{
		name: name,
	}
	f.blockStack = make([]map[string]FloObject, 1)
	f.blockStack[len(f.blockStack)-1] = make(map[string]FloObject, 1)
	vm.callStack = append(vm.callStack, f)

}

func (vm *VM) funcReturn() {

	// f.blockStack[len(f.blockStack)-1] = make(map[string]FloObject, 1)
	vm.callStack = vm.callStack[:len(vm.callStack)-1]

}

func (vm *VM) readVar(name FloString, currentFrame *Frame) FloObject {

	var found bool

	endOfBlockStack := len(currentFrame.blockStack) - 1
	scope := endOfBlockStack
	for j := len(vm.callStack) - 1; j >= 0; j-- {
		for i := scope; i >= 0; i-- {
			if value, ok := vm.callStack[j].blockStack[i][string(name)]; ok {
				found = true
				return value
			}
		}
	}
	if !found {
		panic(nameNotDefinedError(string(name)))
	}
	return nil

	// if value, ok := currentFrame.blockStack[endOfBlockStack][string(name)]; ok {
	// 	return value
	// }
	// panic(nameNotDefinedError(string(name)))

}

func (vm *VM) storeVar(name FloString, val FloObject, currentFrame *Frame) {

	var found bool
	endOfBlockStack := len(currentFrame.blockStack) - 1
	scope := endOfBlockStack
	for i := scope; i >= 0; i-- {
		if _, ok := currentFrame.blockStack[i][string(name)]; ok {
			found = true
			currentFrame.blockStack[i][string(name)] = val
		}
	}

	if !found {
		currentFrame.blockStack[scope][string(name)] = val
	}

}

func (vm *VM) newBlock(currentFrame *Frame) {
	currentFrame.blockStack = append(currentFrame.blockStack, make(map[string]FloObject, 1))
}

func (vm *VM) remBlock(currentFrame *Frame) {
	currentFrame.blockStack = currentFrame.blockStack[:len(currentFrame.blockStack)-1]
}

func (vm *VM) Debug(obj Object) {

	currentFrame := vm.tos()
	fmt.Println("Instructions (text)    :")
	vm.DecodeObjectInstructions(obj)
	fmt.Println("Instructions (bytecode):", obj.Instructions)
	fmt.Println("Contstants             :", obj.Constants)
	fmt.Println("Names                  :", obj.Names)
	fmt.Println("Env                    :", currentFrame.blockStack)

}

// Init the virtual machine
func (vm *VM) Init() {

	vm.funcCall("main")

}

func (vm *VM) getParam(extended_arg int, next_instruction byte, currentFrame *Frame) uint32 {

	byte_array := make([]byte, 4)
	byte_array[3] = next_instruction
	i := 2
	for extended_arg > 0 {
		byte_array[i] = byte(currentFrame.dataStack.Pop().(FloInteger))
		i--
		extended_arg--
	}
	param := binary.BigEndian.Uint32(byte_array)

	return param

}

// Run executes the VM
func (vm *VM) Run(function FloCallable, param_vals FloList) FloObject {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("-- traceback --")
			for i, _ := range vm.callStack {
				frame := vm.callStack[i]
				fmt.Println(frame.name)
			}
			fmt.Println(r)
			// fmt.Println(debug.Stack())
		}
	}()

	instructionCount := len(function.Object.Instructions)
	instructions := function.Object.Instructions
	currentFrame := vm.tos()

	// for i, arg := range function.Args {
	// 	currentFrame.blockStack[0][arg.String()] = param_vals[i]
	// }
	for i, arg := range function.Args {
		// currentFrame.blockStack[0][arg.String()] = param_vals[i]
		vm.storeVar(arg.(FloString), param_vals[i], currentFrame)
	}

	// vm.storeVar(function.Name, function, currentFrame)
	// currentFrame.blockStack[0][function.Name.String()] = function

	vm.Debug(function.Object)
	// fmt.Println("---------")

	var extended_arg int

	for i := 0; i < instructionCount; i += 2 {
		switch instructions[i] {
		case LOAD_CONST:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			currentFrame.dataStack.Push(function.Object.Constants[param])
		case LOAD_NAME:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			x := vm.readVar(function.Object.Names[param], currentFrame)
			currentFrame.dataStack.Push(x)
		case STORE_NAME:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			tos := currentFrame.dataStack.Pop()
			vm.storeVar(function.Object.Names[param], tos, currentFrame)
		case EXTENDED_ARG:
			extended_arg++
			currentFrame.dataStack.Push(FloInteger(instructions[i+1]))
		case PUSH_BLOCK:
			vm.newBlock(currentFrame)
		case POP_BLOCK:
			vm.remBlock(currentFrame)
		case BUILD_LIST:
			list := make([]FloObject, 0)
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			var y uint32
			// Stacks reverse things, this unreverses the list items
			for y = 0; y < param; y++ {
				list = append(list, FloInteger(0))
				copy(list[1:], list)
				list[0] = currentFrame.dataStack.Pop()
			}
			currentFrame.dataStack.Push(FloList(list))
		case GET_ITEM:
			list_index := currentFrame.dataStack.Pop()
			list := currentFrame.dataStack.Pop()
			currentFrame.dataStack.Push((list.(FloList))[list_index.(FloInteger)])
		case SETUP_FUNCTION:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			x := function.Object.Names[param]
			currentFrame.dataStack.Push(x)
			continue
		case SETUP_PARAMS:
			list := make([]FloObject, 0)
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			var y uint32
			// Stacks reverse things, this unreverses the list items
			for y = 0; y < param; y++ {
				list = append(list, FloInteger(0))
				copy(list[1:], list)
				list[0] = currentFrame.dataStack.Pop()
			}
			currentFrame.dataStack.Push(FloList(list))
			continue
		case SETUP_BODY:
			j := 0
			i += 2
			body_start := i
			// Instruction pointer, points to start of function block
			for ok := true; ok; ok = j > 0 {
				if instructions[i] == PUSH_BLOCK {
					j++
				} else if instructions[i] == POP_BLOCK {
					j--
				}
				i += 2
			}
			i -= 2
			body_end := i
			currentFrame.dataStack.Push(FloInteger(body_end - body_start))
			continue
		case MAKE_FUNCTION:
			bodySize := int(currentFrame.dataStack.Pop().(FloInteger))
			params := currentFrame.dataStack.Pop().(FloList)
			name := currentFrame.dataStack.Pop().(FloString)
			functionBody := make([]byte, 0)
			functionBody = append(functionBody, instructions[i-bodySize:i]...)
			object := Object{
				ConstantCount: function.Object.ConstantCount,
				Constants:     make([]FloObject, len(function.Object.Constants)),
				NameCount:     function.Object.NameCount,
				Names:         make([]FloString, len(function.Object.Names)),
				Instructions:  functionBody,
			}
			f := FloCallable{
				Name:   name,
				Args:   params,
				Object: object,
			}
			copy(object.Constants, function.Object.Constants)
			copy(object.Names, function.Object.Names)
			currentFrame.dataStack.Push(f)
			continue
		case CALL_FUNCTION:
			// Check arity
			f := currentFrame.dataStack.Pop().(FloCallable)
			params := int(instructions[i+1])

			if len(f.Args) != params {
				panic(callableArgsTypeError(f, params))
			}

			param_vals := []FloObject{}
			for range f.Args {
				param_vals = append([]FloObject{currentFrame.dataStack.Pop()}, param_vals...)
			}

			vm.funcCall(f.Name.String())
			currentFrame = vm.tos()

			for i, arg := range f.Args {
				// fmt.Println(arg, param_vals[i])
				vm.storeVar(arg.(FloString), param_vals[i], currentFrame)
			}

			function_return := vm.Run(f, param_vals)
			vm.funcReturn()
			currentFrame = vm.tos()
			if function_return != nil {
				currentFrame.dataStack.Push(function_return)
			} else {
				currentFrame.dataStack.Push(FloNil{})
			}
			// fmt.Println(currentFrame.dataStack)
			continue
		case RETURN:
			// for _, call := range vm.callStack {
			// 	fmt.Println(call)
			// }
			// fmt.Println(currentFrame.dataStack)
			return currentFrame.dataStack.Pop()
		case BINARY_ADD:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Add(tos1, tos)
			currentFrame.dataStack.Push(x)
		case BINARY_SUB:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Sub(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_DIV:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Div(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_MUL:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Mul(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_POW:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Pow(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_MOD:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Mod(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_LSHIFT:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_LShift(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_RSHIFT:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_RShift(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_AND:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_And(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_XOR:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_XOR(tos1, tos)
			currentFrame.dataStack.Push(x)

		case BINARY_OR:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := Binary_Or(tos1, tos)
			currentFrame.dataStack.Push(x)

		case UNARY_NOT:
			tos := currentFrame.dataStack.Pop()
			x := Unary_Not(tos)
			currentFrame.dataStack.Push(x)

		case UNARY_ADD:
			tos := currentFrame.dataStack.Pop()
			x := Unary_Add(tos)
			currentFrame.dataStack.Push(x)

		case UNARY_SUB:
			tos := currentFrame.dataStack.Pop()
			x := Unary_Sub(tos)
			currentFrame.dataStack.Push(x)
		case LOGICAL_OR:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := LogicalOr(tos1, tos)
			currentFrame.dataStack.Push(x)
		case LOGICAL_AND:
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			x := LogicalAnd(tos1, tos)
			currentFrame.dataStack.Push(x)
		case COMPARE:
			param := instructions[i+1]
			tos := currentFrame.dataStack.Pop()
			tos1 := currentFrame.dataStack.Pop()
			// fmt.Println("tos", tos, "tos1", tos1)
			x := Compare(tos1, tos, param)
			currentFrame.dataStack.Push(x)

		case POP_JUMP_IF_FALSE:

			param := instructions[i+1]
			tos := currentFrame.dataStack.Pop()
			if tos == FloBool(false) {
				i += (int(param) - 2)
				continue
			}

		// case POP_JUMP_IF_TRUE:

		// 	param := instructions[i+1]
		// 	tos := currentFrame.dataStack.Pop()
		// 	if tos == FloBool(true) {
		// 		i += int(param) - 2
		// 		continue
		// 	}
		case JUMP_BACK:
			param := instructions[i+1]
			i -= (int(param) + 2)

		// case JUMP_IF_FALSE:

		// 	param := instructions[i+1]
		// 	tos := currentFrame.dataStack.Pop()
		// 	if tos == FloBool(false) {
		// 		i += int(param) + 2
		// 	}
		// 	currentFrame.dataStack.Push(tos)

		case JUMP:
			param := instructions[i+1]
			i += int(param) - 2
			continue

		case PRINT:
			Print(currentFrame.dataStack.Pop())

		case NOP:
			continue

		default:
			panic("Unknown instruction")
		}
	}
	return FloNil{}

}

func (vm *VM) DecodeObjectInstructions(obj Object) {

	instructionCount := len(obj.Instructions)
	instructions := obj.Instructions
	currentFrame := vm.tos()
	var extended_arg int

	for i := 0; i < instructionCount; i += 2 {

		switch instructions[i] {
		case LOAD_CONST:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			fmt.Printf("%d LOAD_CONST %s (%d)\n", i, obj.Constants[param], param)
		case LOAD_NAME:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			fmt.Printf("%d LOAD_NAME %s (%d)\n", i, obj.Names[param], param)
		case STORE_NAME:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			fmt.Printf("%d STORE_NAME %s (%d)\n", i, obj.Names[param], param)
		case EXTENDED_ARG:
			fmt.Printf("%d EXTENDED_ARG %d\n", i, instructions[i+1])
		case PUSH_BLOCK:
			fmt.Printf("%d EXTENDED_ARG %d\n", i, instructions[i+1])
		case POP_BLOCK:
			fmt.Printf("%d EXTENDED_ARG %d\n", i, instructions[i+1])
		case BUILD_LIST:
			param := vm.getParam(extended_arg, instructions[i+1], currentFrame)
			fmt.Printf("%d BUILD_LIST %d\n", i, param)
		case GET_ITEM:
			fmt.Printf("%d GET_ITEM %d\n", i, instructions[i+1])
		case SETUP_FUNCTION:
			fmt.Printf("%d SETUP_FUNCTION %d\n", i, instructions[i+1])
		case SETUP_PARAMS:
			fmt.Printf("%d SETUP_PARAMS %d\n", i, instructions[i+1])
		case SETUP_BODY:
			fmt.Printf("%d SETUP_BODY %d\n", i, instructions[i+1])
		case MAKE_FUNCTION:
			fmt.Printf("%d MAKE_FUNCTION %d\n", i, instructions[i+1])
		case CALL_FUNCTION:
			fmt.Printf("%d CALL_FUNCTION %d\n", i, instructions[i+1])
		case RETURN:
			fmt.Printf("%d RETURN %d\n", i, instructions[i+1])
		case BINARY_ADD:
			fmt.Printf("%d BINARY_ADD %d\n", i, instructions[i+1])
		case BINARY_SUB:
			fmt.Printf("%d BINARY_SUB %d\n", i, instructions[i+1])
		case BINARY_DIV:
			fmt.Printf("%d BINARY_DIV %d\n", i, instructions[i+1])
		case BINARY_MUL:
			fmt.Printf("%d BINARY_MUL %d\n", i, instructions[i+1])
		case BINARY_POW:
			fmt.Printf("%d BINARY_POW %d\n", i, instructions[i+1])
		case BINARY_MOD:
			fmt.Printf("%d BINARY_MOD %d\n", i, instructions[i+1])
		case BINARY_LSHIFT:
			fmt.Printf("%d BINARY_LSHIFT %d\n", i, instructions[i+1])
		case BINARY_RSHIFT:
			fmt.Printf("%d BINARY_RSHIFT %d\n", i, instructions[i+1])
		case BINARY_AND:
			fmt.Printf("%d BINARY_AND %d\n", i, instructions[i+1])
		case BINARY_XOR:
			fmt.Printf("%d BINARY_XOR %d\n", i, instructions[i+1])
		case BINARY_OR:
			fmt.Printf("%d BINARY_OR %d\n", i, instructions[i+1])
		case UNARY_NOT:
			fmt.Printf("%d UNARY_NOT %d\n", i, instructions[i+1])
		case UNARY_ADD:
			fmt.Printf("%d UNARY_ADD %d\n", i, instructions[i+1])
		case UNARY_SUB:
			fmt.Printf("%d UNARY_SUB %d\n", i, instructions[i+1])
		case LOGICAL_OR:
			fmt.Printf("%d LOGICAL_OR %d\n", i, instructions[i+1])
		case LOGICAL_AND:
			fmt.Printf("%d LOGICAL_AND %d\n", i, instructions[i+1])
		case COMPARE:
			fmt.Printf("%d COMPARE %d\n", i, instructions[i+1])
		case POP_JUMP_IF_FALSE:
			fmt.Printf("%d POP_JUMP_IF_FALSE %d\n", i, instructions[i+1])
		case POP_JUMP_IF_TRUE:
			fmt.Printf("%d POP_JUMP_IF_TRUE %d\n", i, instructions[i+1])
		case JUMP_BACK:
			fmt.Printf("%d JUMP_BACK %d\n", i, instructions[i+1])
		case JUMP_IF_FALSE:
			fmt.Printf("%d JUMP_IF_FALSE %d\n", i, instructions[i+1])
		case JUMP:
			fmt.Printf("%d JUMP %d\n", i, instructions[i+1])
		case PRINT:
			fmt.Printf("%d PRINT %d\n", i, instructions[i+1])
		case NOP:
			fmt.Printf("%d NOP %d\n", i, instructions[i+1])
		default:
			fmt.Println(i, "UNKNOWN")
		}

	}

}

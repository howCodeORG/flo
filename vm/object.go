package vm

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"os"
)

type Object struct {
	ConstantCount uint32
	Constants     []FloObject
	NameCount     uint32
	Names         []FloString
	Instructions  []byte
	ObjectCode    []byte
}

func (o *Object) getName(name FloString) (uint32, bool) {
	var nameIndex uint32
	var found bool
	for i, n := range o.Names {
		if n == name {
			nameIndex = uint32(i)
			found = true
		}
	}
	return nameIndex, found
}

func (o *Object) insertExtendedArg(val uint32) []byte {
	byte_array := make([]byte, 4)
	binary.BigEndian.PutUint32(byte_array, val)
	for i := 0; i < 3; i++ {
		if byte_array[i] != 0 {
			o.Instructions = append(o.Instructions, EXTENDED_ARG)
			o.Instructions = append(o.Instructions, byte_array[i])
		}
	}
	return byte_array
}

func (o *Object) Encode() []byte {

	var x FloInteger
	var y FloFloat
	var z FloBool
	gob.Register(x)
	gob.Register(y)
	gob.Register(z)
	gob.Register(FloList{})

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(o); err != nil {
		panic(err)
	}

	o.ObjectCode = buf.Bytes()

	return buf.Bytes()

}

func (o *Object) Decode() Object {

	var obj Object
	buf := &bytes.Buffer{}
	buf.Write(o.ObjectCode)

	if err := gob.NewDecoder(buf).Decode(&obj); err != nil {
		panic(obj)
	}

	return obj

}

func (o *Object) Save(fileName string) {

	// create a file
	dataFile, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(o)

	dataFile.Close()
}

func (o *Object) Load(fileName string) {

	var data Object

	// open data file
	dataFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	o.ConstantCount = data.ConstantCount
	o.Constants = data.Constants
	o.NameCount = data.NameCount
	o.Names = data.Names
	o.Instructions = data.Instructions
	o.ObjectCode = data.ObjectCode

	dataFile.Close()

}

func (o *Object) Compile_Load_Const(x FloObject) {
	var constantIndex uint32
	var found bool
	for i, c := range o.Constants {
		if c == x {
			constantIndex = uint32(i)
			found = true
		}
	}
	if found == false {
		constantIndex = o.ConstantCount
		o.Constants = append(o.Constants, x)
		o.ConstantCount++
	}

	byte_array := o.insertExtendedArg(constantIndex)

	o.Instructions = append(o.Instructions, LOAD_CONST)
	o.Instructions = append(o.Instructions, byte_array[3])
}

func (o *Object) Compile_Load_Name(name FloString) {

	nameIndex, found := o.getName(name)

	if found == false {
		nameIndex = o.NameCount
		o.Names = append(o.Names, name)
		o.NameCount++
	}

	byte_array := o.insertExtendedArg(nameIndex)

	o.Instructions = append(o.Instructions, LOAD_NAME)
	o.Instructions = append(o.Instructions, byte_array[3])

	// o.Instructions = append(o.Instructions, LOAD_NAME)
	// o.Instructions = append(o.Instructions, byte(nameIndex))
}

func (o *Object) Compile_Store_Name(name FloString) {

	nameIndex, found := o.getName(name)

	if !found {
		nameIndex = o.NameCount
		o.Names = append(o.Names, name)
		o.NameCount++
	}

	byte_array := o.insertExtendedArg(nameIndex)

	o.Instructions = append(o.Instructions, STORE_NAME)
	o.Instructions = append(o.Instructions, byte_array[3])

	// o.Instructions = append(o.Instructions, STORE_NAME)
	// o.Instructions = append(o.Instructions, byte(nameIndex))
}

func (o *Object) Compile_Push_Block() {

	o.Instructions = append(o.Instructions, PUSH_BLOCK)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Pop_Block() {

	o.Instructions = append(o.Instructions, POP_BLOCK)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Build_List(size uint32) {

	o.Instructions = append(o.Instructions, BUILD_LIST)
	o.Instructions = append(o.Instructions, byte(size))

}

func (o *Object) Compile_Get_Item() {

	o.Instructions = append(o.Instructions, GET_ITEM)
	o.Instructions = append(o.Instructions, byte(0))

}

func (o *Object) Compile_Setup_Function(name FloString) {

	nameIndex, found := o.getName(name)

	if !found {
		nameIndex = o.NameCount
		o.Names = append(o.Names, name)
		o.NameCount++
	}

	byte_array := o.insertExtendedArg(nameIndex)

	o.Instructions = append(o.Instructions, SETUP_FUNCTION)
	o.Instructions = append(o.Instructions, byte_array[3])

}

func (o *Object) Compile_Setup_Params(num FloInteger) {

	o.Instructions = append(o.Instructions, SETUP_PARAMS)
	o.Instructions = append(o.Instructions, byte(num))

}

func (o *Object) Compile_Setup_Body() {

	o.Instructions = append(o.Instructions, SETUP_BODY)
	o.Instructions = append(o.Instructions, byte(0))

}

func (o *Object) Compile_Make_Function() {

	o.Instructions = append(o.Instructions, MAKE_FUNCTION)
	o.Instructions = append(o.Instructions, byte(0))

}

func (o *Object) Compile_Call_Function(args int) {

	o.Instructions = append(o.Instructions, CALL_FUNCTION)
	o.Instructions = append(o.Instructions, byte(args))

}

func (o *Object) Compile_Return() {

	o.Instructions = append(o.Instructions, RETURN)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Add() {

	o.Instructions = append(o.Instructions, BINARY_ADD)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Sub() {

	o.Instructions = append(o.Instructions, BINARY_SUB)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Mul() {

	o.Instructions = append(o.Instructions, BINARY_MUL)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Div() {

	o.Instructions = append(o.Instructions, BINARY_DIV)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Pow() {

	o.Instructions = append(o.Instructions, BINARY_POW)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Mod() {

	o.Instructions = append(o.Instructions, BINARY_MOD)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_LShift() {

	o.Instructions = append(o.Instructions, BINARY_LSHIFT)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_RShift() {

	o.Instructions = append(o.Instructions, BINARY_RSHIFT)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_And() {

	o.Instructions = append(o.Instructions, BINARY_AND)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_XOR() {

	o.Instructions = append(o.Instructions, BINARY_XOR)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Binary_Or() {

	o.Instructions = append(o.Instructions, BINARY_OR)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Unary_Not() {

	o.Instructions = append(o.Instructions, UNARY_NOT)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Unary_Add() {

	o.Instructions = append(o.Instructions, UNARY_ADD)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Unary_Sub() {

	o.Instructions = append(o.Instructions, UNARY_SUB)
	o.Instructions = append(o.Instructions, 0)

}

func (o *Object) Compile_Unary_Inc(name FloString) {

	o.Compile_Load_Name(name)
	o.Compile_Load_Const(FloInteger(1))
	o.Compile_Binary_Add()
	o.Compile_Store_Name(name)

}

func (o *Object) Compile_Unary_Dec(name FloString) {

	o.Compile_Load_Name(name)
	o.Compile_Load_Const(FloInteger(1))
	o.Compile_Binary_Sub()
	o.Compile_Store_Name(name)

}

func (o *Object) Compile_Compare(op int) {

	o.Instructions = append(o.Instructions, COMPARE)
	o.Instructions = append(o.Instructions, byte(op))

}

func (o *Object) Compile_Pop_Jump_If_False(nop_start int, jump_by int) {

	o.Instructions[nop_start] = POP_JUMP_IF_FALSE
	o.Instructions[nop_start+1] = byte(jump_by)

	// o.Instructions = append(o.Instructions, POP_JUMP_IF_FALSE)
	// o.Instructions = append(o.Instructions, byte(jump_to))

}

func (o *Object) Compile_Pop_Jump_If_True(jump_by int) {

	o.Instructions = append(o.Instructions, POP_JUMP_IF_TRUE)
	o.Instructions = append(o.Instructions, byte(jump_by))

}

func (o *Object) Compile_Jump_Back(jump_by int) {

	o.Instructions = append(o.Instructions, JUMP_BACK)
	o.Instructions = append(o.Instructions, byte(jump_by))

}

func (o *Object) Compile_Jump_If_False(nop_start int, jump_by int) {

	o.Instructions[nop_start] = JUMP_IF_FALSE
	o.Instructions[nop_start+1] = byte(jump_by)

	// o.Instructions = append(o.Instructions, POP_JUMP_IF_FALSE)
	// o.Instructions = append(o.Instructions, byte(jump_to))

}

func (o *Object) Compile_Jump(nop int, jump_by int) {

	o.Instructions[nop] = JUMP
	o.Instructions[nop+1] = byte(jump_by)

}

func (o *Object) Compile_Nop(num int) {

	for i := 0; i < num; i++ {
		o.Instructions = append(o.Instructions, NOP)
		o.Instructions = append(o.Instructions, byte(0))
	}

}

func (o *Object) Compile_Print() {

	o.Instructions = append(o.Instructions, PRINT)
	o.Instructions = append(o.Instructions, 0)

}

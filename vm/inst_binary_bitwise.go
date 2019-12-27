package vm

// Binary_And implements BINARY_AND
func Binary_And(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) & (right.(FloInteger))
	} else {
		panic(unsupportedTypesError("&", left, right))
	}

	return value

}

// Binary_XOR implements BINARY_XOR
func Binary_XOR(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) ^ (right.(FloInteger))
	} else {
		panic(unsupportedTypesError("^", left, right))
	}

	return value

}

// Binary_Or implements BINARY_OR
func Binary_Or(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) | (right.(FloInteger))
	} else {
		panic(unsupportedTypesError("|", left, right))
	}

	return value

}

package vm

// Binary_LShift implements BINARY_LSHIFT
func Binary_LShift(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) << uint(right.(FloInteger))
	} else {
		panic(unsupportedTypesError("<<", left, right))
	}

	return value

}

// Binary_RShift implements BINARY_RSHIFT
func Binary_RShift(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) >> uint(right.(FloInteger))
	} else {
		panic(unsupportedTypesError(">>", left, right))
	}

	return value

}

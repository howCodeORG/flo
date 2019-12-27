package eval

// BitwiseAnd implements &
func BitwiseAnd(left, right FloObject) FloObject {

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

// BitwiseXOR implements ^
func BitwiseXOR(left, right FloObject) FloObject {

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

// BitwiseOr implements |
func BitwiseOr(left, right FloObject) FloObject {

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

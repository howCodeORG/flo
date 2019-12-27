package eval

// Minus implements subtraction for Flo's primitive types
func Minus(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) - right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) - right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) - FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) - right.(FloFloat)
	} else {
		panic(unsupportedTypesError("-", left, right))
	}

	return value

}

// UnaryMinus implements unary Subtraction for Flo's primitive types
func UnaryMinus(x FloObject) FloObject {

	var value FloObject

	t := x.Type()

	if t == FloINT {
		value = -(x.(FloInteger))
	} else if t == FloFLOAT {
		value = -(x.(FloFloat))
	} else {
		panic(unsupportedUnaryTypeError("-", x))
	}

	return value

}

// UnaryMinusMinus implements unary MinusMinus for Flo's primitive types
func UnaryMinusMinus(x FloObject) FloObject {

	var value FloObject

	t := x.Type()

	if t == FloINT {
		value = FloInteger(x.(FloInteger)) - 1
	} else if t == FloFLOAT {
		value = FloFloat(x.(FloFloat)) - 1
	} else {
		panic(unsupportedUnaryTypeError("--", x))
	}

	return value

}

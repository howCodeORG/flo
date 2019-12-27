package eval

// Plus implements Addition for Flo's primitive types
func Plus(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) + right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) + right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) + FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) + right.(FloFloat)
	} else if t1 == FloSTRING && t2 == FloSTRING {
		value = left.(FloString) + right.(FloString)
	} else {
		panic(unsupportedTypesError("+", left, right))
	}

	return value

}

// UnaryPlus implements unary Addition for Flo's primitive types
func UnaryPlus(x FloObject) FloObject {

	var value FloObject

	t := x.Type()

	if t == FloINT {
		value = +(x.(FloInteger))
	} else if t == FloFLOAT {
		value = +(x.(FloFloat))
	} else {
		panic(unsupportedUnaryTypeError("+", x))
	}

	return value

}

// UnaryPlusPlus implements unary PlusPlus for Flo's primitive types
func UnaryPlusPlus(x FloObject) FloObject {

	var value FloObject

	t := x.Type()

	if t == FloINT {
		value = FloInteger(x.(FloInteger)) + 1
	} else if t == FloFLOAT {
		value = FloFloat(x.(FloFloat)) + 1
	} else {
		panic(unsupportedUnaryTypeError("++", x))
	}

	return value

}

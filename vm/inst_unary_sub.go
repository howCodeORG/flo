package vm

// Unary_Sub implements UNARY_INC
func Unary_Sub(x FloObject) FloObject {

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

// Unary_Dec implements UNARY_DEC
func Unary_Dec(x FloObject) FloObject {

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

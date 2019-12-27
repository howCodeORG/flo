package vm

// Unary_Not implements logical not (UNARY_NOT)
func Unary_Not(x FloObject) FloBool {

	var value FloBool

	t := x.Type()

	if t == FloBOOL {
		value = !(x.(FloBool))
	} else {
		panic(unsupportedUnaryTypeError("not", x))
	}

	return value

}

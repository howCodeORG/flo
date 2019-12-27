package eval

// Not implements logical not for Flo's primitive types
func Not(x FloObject) FloBool {

	var value FloBool

	t := x.Type()

	if t == FloBOOL {
		value = !(x.(FloBool))
	} else {
		panic(unsupportedUnaryTypeError("not", x))
	}

	return value

}

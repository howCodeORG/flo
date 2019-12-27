package eval

// And implements logical and for Flo's primitive types
func And(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloBOOL && t2 == FloBOOL {
		value = left.(FloBool) && right.(FloBool)
	} else {
		panic(unsupportedTypesError("and", left, right))
	}

	return value

}

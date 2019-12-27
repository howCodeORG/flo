package eval

// Or implements logical or for Flo's primitive types
func Or(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloBOOL && t2 == FloBOOL {
		value = left.(FloBool) || right.(FloBool)
	} else {
		panic(unsupportedTypesError("or", left, right))
	}

	return value

}

package vm

// LogicalOr implements LOGICAL_OR
func LogicalOr(left, right FloObject) FloBool {

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

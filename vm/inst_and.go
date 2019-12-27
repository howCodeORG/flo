package vm

// LogicalAnd implements LOGICAL_AND
func LogicalAnd(left, right FloObject) FloObject {

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

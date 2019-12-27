package vm

// BinaryAdd implements BINARY_ADD
func Binary_Add(left, right FloObject) FloObject {

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

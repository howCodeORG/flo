package eval

// LShift implements <<
func LShift(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) << uint(right.(FloInteger))
	} else {
		panic(unsupportedTypesError("<<", left, right))
	}

	return value

}

// RShift implements >>
func RShift(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) >> uint(right.(FloInteger))
	} else {
		panic(unsupportedTypesError(">>", left, right))
	}

	return value

}

package vm

func Compare(left FloObject, right FloObject, op byte) FloBool {
	switch op {
	case 0:
		return LessThan(left, right)
	case 1:
		return LessThanEqual(left, right)
	case 2:
		return GreaterThan(left, right)
	case 3:
		return GreaterThanEqual(left, right)
	case 4:
		return Equals(left, right)
	case 5:
		return NotEquals(left, right)
	default:
		panic(unsupportedTypesError("", left, right))
	}
	return FloBool(false)
}

// LessThan implements <
func LessThan(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) < right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) < right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) < FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) < right.(FloFloat)
	} else {
		panic(unsupportedTypesError("<", left, right))
	}

	return value

}

// LessThanEqual implements <=
func LessThanEqual(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) <= right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) <= right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) <= FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) <= right.(FloFloat)
	} else {
		panic(unsupportedTypesError("<=", left, right))
	}

	return value

}

// GreaterThan implements >
func GreaterThan(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) > right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) > right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) > FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) > right.(FloFloat)
	} else {
		panic(unsupportedTypesError(">", left, right))
	}

	return value

}

// GreaterThanEqual implements >=
func GreaterThanEqual(left, right FloObject) FloBool {

	var value FloBool

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) >= right.(FloInteger)
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(left.(FloInteger)) >= right.(FloFloat)
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = left.(FloFloat) >= FloFloat(right.(FloInteger))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = left.(FloFloat) >= right.(FloFloat)
	} else {
		panic(unsupportedTypesError(">=", left, right))
	}

	return value

}

// Equals implements ==
func Equals(left, right FloObject) FloBool {

	var value FloBool

	value = left == right

	return value

}

// NotEquals implements !=
func NotEquals(left, right FloObject) FloBool {

	var value FloBool

	value = left != right

	return value

}

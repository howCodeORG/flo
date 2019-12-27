package eval

// Mod implements modulo for Flo's primitive types
func Mod(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = left.(FloInteger) % right.(FloInteger)
	} else {
		panic(unsupportedTypesError("%", left, right))
	}

	return value

}

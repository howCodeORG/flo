package vm

import "math"

// Binary_Pow implements BINARY_POWER
func Binary_Pow(left, right FloObject) FloObject {

	var value FloObject

	t1 := left.Type()
	t2 := right.Type()

	if t1 == FloINT && t2 == FloINT {
		value = FloInteger(math.Pow(float64(left.(FloInteger)), float64(right.(FloInteger))))
	} else if t1 == FloINT && t2 == FloFLOAT {
		value = FloFloat(math.Pow(float64(left.(FloInteger)), float64(right.(FloFloat))))
	} else if t1 == FloFLOAT && t2 == FloINT {
		value = FloFloat(math.Pow(float64(left.(FloFloat)), float64(right.(FloInteger))))
	} else if t1 == FloFLOAT && t2 == FloFLOAT {
		value = FloFloat(math.Pow(float64(left.(FloFloat)), float64(right.(FloFloat))))
	} else {
		panic(unsupportedTypesError("**", left, right))
	}

	return value

}

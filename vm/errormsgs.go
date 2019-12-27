package vm

import "strconv"

func unsupportedTypesError(op string, left, right FloObject) string {
	return "RuntimeError: unsupported type(s) for " + op + ": '" + left.TypeString() + "' and '" + right.TypeString() + "'"
}

func unsupportedUnaryTypeError(op string, obj FloObject) string {
	return "RuntimeError: unsupported type for unary " + op + ": '" + obj.TypeString() + "'"
}

func callableArgsTypeError(obj FloCallable, given int) string {
	return "TypeError: func() expects " + strconv.Itoa(len(obj.Args)) + " arguments but " + strconv.Itoa(given) + ""
}

func notCallableError(obj FloObject) string {
	return "RuntimeError: '" + obj.TypeString() + "' type is not callable."
}

func nameNotDefinedError(name string) string {
	return "NameError: name '" + name + "' is not defined"
}

func indexOutOfRangeError() string {
	return "IndexError: list index out of range"
}

func listIndexTypeError(obj FloObject) string {
	return "TypeError: list indicies must be of type 'int', not '" + obj.TypeString() + "'"
}

func listTypeError(obj FloObject) string {
	return "TypeError: '" + obj.TypeString() + "' is not subscriptable"
}

func conditionNotBoolError(obj FloObject) string {
	return "RuntimeError: condition must evaluate to 'bool', not '" + obj.TypeString() + "'"
}

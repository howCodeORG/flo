package eval

import (
	"flo/parser"
	"fmt"
)

const (
	FloINT = iota
	FloFLOAT
	FloBOOL
	FloSTRING
	FloLIST
	FloCALLABLE
)

type FloObject interface {
	String() string
	Type() int
	TypeString() string
}

type Callable interface {
	Call()
}

type FloInteger int

func (x FloInteger) String() string {
	return fmt.Sprintf("%d", x)
}

func (x FloInteger) Type() int {
	return FloINT
}

func (x FloInteger) TypeString() string {
	return "int"
}

type FloFloat float64

func (x FloFloat) String() string {
	return fmt.Sprintf("%f", x)
}

func (x FloFloat) Type() int {
	return FloFLOAT
}

func (x FloFloat) TypeString() string {
	return "float"
}

type FloBool bool

func (x FloBool) String() string {
	return fmt.Sprintf("%t", x)
}

func (x FloBool) Type() int {
	return FloBOOL
}

func (x FloBool) TypeString() string {
	return "bool"
}

type FloString string

func (x FloString) String() string {
	return "'" + string(x) + "'"
}

func (x FloString) Type() int {
	return FloSTRING
}

func (x FloString) TypeString() string {
	return "string"
}

type FloList []FloObject

func (x FloList) String() string {

	var output string
	output = "["

	for _, val := range x {
		if len(x) == 1 {
			output += fmt.Sprintf("%s", val)
		} else {
			output += fmt.Sprintf("%s, ", val)
		}
	}

	if len(x) > 1 {
		output = output[:len(output)-2]
	}
	output += "]"

	return output

}

func (x FloList) Type() int {
	return FloLIST
}

func (x FloList) TypeString() string {
	return "list"
}

type FloCallable struct {
	name      string
	min_scope int
	args      FloList
	body      parser.IBlockContext
}

func (x FloCallable) String() string {
	return "'" + x.name + "'"
}

func (x FloCallable) Type() int {
	return FloSTRING
}

func (x FloCallable) TypeString() string {
	return "callable"
}

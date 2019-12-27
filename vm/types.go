package vm

import (
	"fmt"
)

const (
	FloNIL = iota
	FloINT
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

type FloNil struct{}

func (x FloNil) String() string {
	return fmt.Sprintf("nil")
}

func (x FloNil) Type() int {
	return FloNIL
}

func (x FloNil) TypeString() string {
	return "nil"
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
	Args   FloList
	Name   FloString
	Object Object
}

func (x FloCallable) String() string {
	return fmt.Sprintf("func "+string(x.Name)+"("+x.Args.String()+") at %p", &x.Object)
}

func (x FloCallable) Type() int {
	return FloSTRING
}

func (x FloCallable) TypeString() string {
	return "callable"
}

package object

import (
	"flo/ast"
	"fmt"
)

type Object interface {
	Type() string
	String() string
}

const (
	INTEGER_OBJECT  = "int"
	FLOAT_OBJECT    = "float"
	BOOLEAN_OBJECT  = "bool"
	STRING_OBJECT   = "string"
	FUNCTION_OBJECT = "func"
	LIST_OBJECT     = "list"
	NIL_OBJECT      = "nil"
)

// NIL OBJECT
type Nil struct {
}

func (n *Nil) Type() string {
	return NIL_OBJECT
}

func (n *Nil) String() string {
	return fmt.Sprintf("nil")
}

// INTEGER OBJECT
type Integer struct {
	Value int
}

func (i *Integer) Type() string {
	return INTEGER_OBJECT
}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// STRING OBJECT
type String struct {
	Value string
}

func (s *String) Type() string {
	return STRING_OBJECT
}

func (s *String) String() string {
	return fmt.Sprintf("%s", s.Value)
}

// BOOL OBJECT
type Bool struct {
	Value bool
}

func (b *Bool) Type() string {
	return BOOLEAN_OBJECT
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.Value)
}

// FLOAT OBJECT
type Float struct {
	Value float64
}

func (f *Float) Type() string {
	return FLOAT_OBJECT
}

func (f *Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}

// FUNCTION OBJECT
type Function struct {
	Name       string
	Parameters []ast.Noder
	Code       func(params ...ast.Noder) ast.Noder
	Builtin    bool
}

func (f *Function) Type() string {
	return FUNCTION_OBJECT
}

func (f *Function) String() string {
	params := ""
	if f.Parameters != nil {
		for _, param := range f.Parameters {
			params = params + param.String() + ", "
		}
		if params != "" {
			params = params[:len(params)-2]
		}
	}
	return fmt.Sprintf("%s(%s)", f.Name, params)
}

// LIST OBJECT
type List struct {
	Values []Object
}

func (l *List) Type() string {
	return LIST_OBJECT
}

func (l *List) String() string {
	values := "["
	if l.Values != nil {
		for _, value := range l.Values {
			values = values + value.String() + ", "
		}
		values = values[:len(values)-2]
	}
	values += "]"
	return values
}

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
	VOID_OBJECT     = "void"
)

// VOID OBJECT
type Void struct {
}

func (v *Void) Type() string {
	return VOID_OBJECT
}

func (v *Void) String() string {
	return fmt.Sprintf("void")
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

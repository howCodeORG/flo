package ast

type Noder interface {
	String() string
}

// CURRENT AST NODES
// ast.Node
// ast.List
// ast.Statement
// ast.UnaryOp
// ast.BinaryOp
// ast.IfStatement
// ast.Function
// ast.Block

// ast.Node
type Node struct {
	Type  string
	Value string
}

func (n *Node) String() string {
	return n.Value
}

// ast.List
type List struct {
	Values []Noder
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

// ast.Statement
type Statement struct {
	Value Noder
	Next  Noder
}

func (s *Statement) String() string {
	value := "<nil>"
	next := "<nil>"
	if s.Value != nil {
		value = s.Value.String()
	}
	if s.Next != nil {
		next = s.Next.String()
	}
	return value + " -> " + next
}

// ast.UnaryOp
type UnaryOp struct {
	Operator string
	Value    Noder
}

func (u *UnaryOp) String() string {

	value := "<nil>"

	if u.Value != nil {
		value = u.Value.String()
	}

	return "( " + u.Operator + " " + value + " )"
}

// ast.BinaryOp
type BinaryOp struct {
	Operator string
	Left     Noder
	Right    Noder
}

func (b *BinaryOp) String() string {

	left := "<nil>"
	right := "<nil>"

	if b.Left != nil {
		left = b.Left.String()
	}
	if b.Right != nil {
		right = b.Right.String()
	}

	return "( " + b.Operator + " " + left + " " + right + " )"
}

// ast.IfStatement
type IfStatement struct {
	Type      string
	Condition Noder
	If        Noder
	Else      Noder
}

func (t *IfStatement) String() string {

	condition := "<nil>"
	if_part := "<nil>"
	else_part := "<nil>"

	if t.Condition != nil {
		condition = t.Condition.String()
	}
	if t.If != nil {
		if_part = t.If.String()
	}
	if t.Else != nil {
		else_part = t.Else.String()
	}

	return "( " + t.Type + ", " + condition + ", " + if_part + ", " + else_part + " )"
}

// ast.ForStatement
type ForStatement struct {
	Condition Noder
	Init      Noder
	Post      Noder
	Body      Noder
}

func (f *ForStatement) String() string {

	condition := ""
	init := ""
	post := ""
	body := ""

	if f.Condition != nil {
		condition = f.Condition.String()
	}
	if f.Init != nil {
		init = f.Init.String()
	}
	if f.Post != nil {
		post = f.Post.String()
	}
	if f.Body != nil {
		body = f.Body.String()
	}

	return "( " + condition + ", Init: " + init + ", Post: " + post + ", Body: " + body + " )"
}

// ast.Function
type Function struct {
	Name       string
	Parameters List
	Body       Noder
}

func (f *Function) String() string {

	body := "<nil>"

	if f.Body != nil {
		body = f.Body.String()
	}

	return "( " + f.Name + " " + f.Parameters.String() + " " + body + " )"
}

// ast.Block
type Block struct {
	Value Noder
}

func (b *Block) String() string {
	v := ""
	if b.Value != nil {
		v = b.Value.String()
	}
	return "{ " + v + " }"
}

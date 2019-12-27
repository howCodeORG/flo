package parser

import (
	"flo/ast"
	"flo/lexer"
	"flo/token"
	// "fmt"
	"strconv"
)

type Parser struct {
	l              *lexer.Lexer
	errors         []string
	needToSync     bool
	currentToken   token.Token
	previousToken  token.Token
	lookaheadToken token.Token
	anonFuncCount  int
}

func New(l *lexer.Lexer) *Parser {

	p := &Parser{
		l:             l,
		errors:        []string{},
		needToSync:    false,
		anonFuncCount: 0,
	}
	p.nextToken()
	p.nextToken()
	return p

}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) syntaxError(s string) string {
	return s + " at line " + strconv.Itoa(p.previousToken.Line) + ", col " + strconv.Itoa(p.previousToken.Column)
}

func (p *Parser) FoundError(e string) {
	p.errors = append(p.errors, e)
	p.synchronize()
}

func (p *Parser) synchronize() {
	for p.currentToken.Type != token.SEMICOLON &&
		// p.currentToken.Type != token.ILLEGAL &&
		p.currentToken.Type != token.RBRACE &&
		// p.currentToken.Type != token.K_FUNC &&
		p.currentToken.Type != token.EOF {
		// fmt.Println(p.currentToken)
		p.nextToken()
	}
	p.needToSync = false
}

func (p *Parser) nextToken() token.Token {

	p.previousToken = p.currentToken
	p.currentToken = p.lookaheadToken
	p.lookaheadToken = p.l.NextToken()

	return p.previousToken

}

func (p *Parser) peek() token.Token {
	return p.lookaheadToken
}

func (p *Parser) foundError(errMsg string) bool {
	p.FoundError(errMsg)
	return true
}

func (p *Parser) expectedToken(t []string, errMsg string) bool {
	err := true
	for _, tok := range t {
		if p.currentToken.Type == tok || (p.currentToken.Value == "" && p.currentToken.Type != token.EOF) {
			err = false
			break
		}
	}
	if err == true {

		p.foundError(errMsg)
		p.needToSync = true

		return false
	}
	return true
}

func (p *Parser) unexpectedToken(t []string, errMsg string) bool {
	err := false
	for _, tok := range t {
		if p.currentToken.Type == tok {
			err = true
			break
		}
	}
	if err == true {

		p.foundError(errMsg)
		p.needToSync = true

		return false
	}
	return true
}

func (p *Parser) ParseProgram() ast.Noder {
	return p.program()
}

// program : statement ( ';' program )*
//         | block ( ';' program )*
//         | EOF
func (p *Parser) program() ast.Noder {

	pro := &ast.Statement{}

	if p.currentToken.Type == token.EOF {
		p.nextToken()
		return &ast.Node{Type: "NIL", Value: "nil"}
	}

	if p.currentToken.Type == token.LBRACE {

		pro.Value = p.block()
		for p.currentToken.Type == token.SEMICOLON {
			p.nextToken()
			pro.Next = p.program()
			if pro.Next != nil && pro.Next.String() != "nil" {
				pro = &ast.Statement{Value: pro}
			} else {
				pro.Next = nil
			}
		}

	} else {

		pro.Value = p.statement()
		p.unexpectedToken([]string{token.ILLEGAL, token.RPAREN}, p.syntaxError("syntax error"))
		for p.currentToken.Type == token.SEMICOLON {
			p.nextToken()
			pro.Next = p.program()
			if pro.Next != nil && pro.Next.String() != "nil" {
				pro = &ast.Statement{Value: pro}
			} else {
				pro.Next = nil
			}

		}
	}

	if pro.Next == nil {
		return pro.Value
	}

	return pro

}

// block : '{' program '}'
func (p *Parser) block() ast.Noder {

	block := &ast.Block{}

	p.expectedToken([]string{token.LBRACE}, p.syntaxError("expected {"))
	p.nextToken()

	prog := p.program()

	p.expectedToken([]string{token.RBRACE}, p.syntaxError("expected }"))
	p.nextToken()

	block.Value = prog

	return block
}

// statement                  : assignment
//                            | static_variable_assignment
//                            | if_statement
//                            | for_statement
//                            | func_declaration
//                            | print_statement
//							  | return_statement
//							  | break_statement
//                            | expression
func (p *Parser) statement() ast.Noder {
	if p.peek().Type == token.ASSIGN || p.peek().Type == token.COMMA {
		return p.assignment()
	} else if p.currentToken.Type == token.K_FUNC && p.peek().Type != token.LPAREN {
		return p.funcDeclaration()
	} else if p.currentToken.Type == token.K_IF {
		return p.ifStatement()
	} else if p.currentToken.Type == token.K_FOR {
		return p.forStatement()
	} else if p.currentToken.Type == token.K_PRINT {
		return p.printStmt()
	} else if p.currentToken.Type == token.K_RETURN {
		return p.returnStmt()
	} else if p.currentToken.Type == token.K_BREAK {
		return p.breakStmt()
	} else {
		return p.expression()
	}
}

// if_statement               : IF condition block
//							  : IF condition block else if_statement
//                            | IF condition block else block
func (p *Parser) ifStatement() ast.Noder {

	if_stmt := &ast.IfStatement{}
	if_stmt.Type = "if"

	p.nextToken()

	if_stmt.Condition = p.expression()

	p.expectedToken([]string{token.LBRACE}, p.syntaxError("expected '{'"))
	if_stmt.If = p.block()

	if p.currentToken.Type == token.K_ELSE {
		p.nextToken()
		if p.currentToken.Type == token.K_IF {
			if_stmt.Else = p.ifStatement()
		} else {
			if_stmt.Else = p.block()
		}
	}

	return if_stmt
}

// for_statement : FOR (expression) block
//               | FOR (init ';' expression ';' post) block
//               | FOR block
func (p *Parser) forStatement() ast.Noder {

	for_stmt := &ast.ForStatement{}

	p.nextToken()

	for_stmt.Init = p.forInit()
	if for_stmt.Init == nil {
		if p.currentToken.Type == token.LBRACE {
			// Rule 3
			for_stmt.Condition = &ast.Node{Type: token.BOOL, Value: "true"}
			for_stmt.Body = p.block()

		} else {
			// Rule 1
			for_stmt.Condition = p.expression()
			if p.currentToken.Type == token.SEMICOLON {
				p.nextToken()
			}
			p.expectedToken([]string{token.LBRACE}, p.syntaxError("expecting '{' after expression"))
			for_stmt.Body = p.block()
		}
	} else {
		// Rule 2
		p.expectedToken([]string{token.SEMICOLON}, p.syntaxError("expecting ';' after init"))
		p.nextToken()

		for_stmt.Condition = p.expression()
		p.expectedToken([]string{token.SEMICOLON}, p.syntaxError("expecting ';' after condition"))
		p.nextToken()

		for_stmt.Post = p.forPost()
		if p.currentToken.Type == token.SEMICOLON {
			p.nextToken()
		}
		p.expectedToken([]string{token.LBRACE}, p.syntaxError("expecting '{' after post-condition"))

		for_stmt.Body = p.block()
	}

	return for_stmt

}

// init : assignment
func (p *Parser) forInit() ast.Noder {
	if p.currentToken.Type == token.IDENT && p.peek().Type == token.ASSIGN {
		return p.assignment()
	}
	return nil
}

// post : assignment
func (p *Parser) forPost() ast.Noder {
	return p.assignment()
}

// return_stmt : Return expr
func (p *Parser) returnStmt() ast.Noder {

	p.nextToken()
	value := p.expression()
	// p.nextToken()
	return &ast.UnaryOp{Operator: "return", Value: value}

}

// break_stmt : Break
func (p *Parser) breakStmt() ast.Noder {

	p.nextToken()
	return &ast.UnaryOp{Operator: "break", Value: nil}

}

// printStmt : Print expr
func (p *Parser) printStmt() ast.Noder {

	p.nextToken()
	value := p.expression()
	// p.nextToken()
	return &ast.UnaryOp{Operator: "print", Value: value}

}

// funcDeclaration : Func Identifier '(' parameterList? ')' block
func (p *Parser) funcDeclaration() ast.Noder {

	p.nextToken()
	p.expectedToken([]string{token.IDENT}, p.syntaxError("expected identifier"))
	function := &ast.Function{}
	function.Name = p.nextToken().Value

	p.expectedToken([]string{token.LPAREN}, p.syntaxError("expected ("))
	p.nextToken()

	// parameterList?
	if p.currentToken.Type != token.RPAREN {
		function.Parameters = *(p.identifierList()).(*ast.List)
	}

	p.expectedToken([]string{token.RPAREN}, p.syntaxError("expected )"))
	p.nextToken()

	function.Body = p.block()

	return function

}

// expressionList : expression ( ',' expression )*
func (p *Parser) expressionList() ast.Noder {

	list := &ast.List{}

	// expression
	list.Values = append(list.Values, p.expression())

	if p.currentToken.Type == token.RBRACK {
		p.nextToken()
	}

	// ( ',' expression )*
	for p.currentToken.Type == token.COMMA {
		p.nextToken()
		list.Values = append(list.Values, p.expression())
	}

	return list

}

// assignment : identifierList '=' expresssionList
func (p *Parser) assignment() ast.Noder {

	assignment := &ast.BinaryOp{}
	p.expectedToken([]string{token.IDENT}, p.syntaxError("left-hand side of assignment should be identifier"))
	assignment.Left = p.identifierList()
	assignment.Operator = p.nextToken().Value
	assignment.Right = p.expressionList()

	if len((assignment.Left.(*ast.List)).Values) != len((assignment.Right.(*ast.List)).Values) {
		p.foundError(p.syntaxError("unbalanced assignment"))
	}

	return assignment

}

func (p *Parser) expression() ast.Noder {
	return p.equality_expression()
}

// equality_expression : relational_expression ( ( '==', '!=' ) relational_expression)*
func (p *Parser) equality_expression() ast.Noder {

	expr := &ast.BinaryOp{}
	expr.Left = p.relational_expression()

	// p.expectedToken([]string{
	// 	token.EQEQ,
	// 	token.NTEQ,
	// 	token.SEMICOLON,
	// 	token.COMMA,
	// 	token.RBRACE,
	// 	token.EOF}, // Don't show syntax errors if we find a semicolon or a comma or a '}' or EOF
	// 	p.syntaxError("found '"+p.currentToken.Value+"'"))

	for p.currentToken.Type == token.EQEQ ||
		p.currentToken.Type == token.NTEQ {

		expr.Operator = p.nextToken().Value
		expr.Right = p.relational_expression()

		expr = &ast.BinaryOp{Left: expr}

	}

	return expr.Left

}

// relational_expression : additive_expression ( ( '<' | '>' | '<=' | '>=' ) additive_expression)*
func (p *Parser) relational_expression() ast.Noder {

	expr := &ast.BinaryOp{}
	expr.Left = p.additive_expression()

	for p.currentToken.Type == token.LT ||
		p.currentToken.Type == token.GT ||
		p.currentToken.Type == token.LTEQ ||
		p.currentToken.Type == token.GTEQ {

		expr.Operator = p.nextToken().Value
		expr.Right = p.additive_expression()

		expr = &ast.BinaryOp{Left: expr}

	}

	return expr.Left

}

// additive_expression : term ( ( '+' | '-' ) term)*
func (p *Parser) additive_expression() ast.Noder {

	expr := &ast.BinaryOp{}
	expr.Left = p.term()

	for p.currentToken.Type == token.PLUS ||
		p.currentToken.Type == token.MINUS {

		expr.Operator = p.nextToken().Value
		expr.Right = p.term()

		expr = &ast.BinaryOp{Left: expr, Right: nil}

	}

	return expr.Left

}

// term : factor ( ( '*' | '/' | '%' ) factor)*
func (p *Parser) term() ast.Noder {

	expr := &ast.BinaryOp{}
	expr.Left = p.factor()

	for p.currentToken.Type == token.MULTIPLY ||
		p.currentToken.Type == token.DIVIDE ||
		p.currentToken.Type == token.MODULO {

		expr.Operator = p.nextToken().Value
		expr.Right = p.factor()

		expr = &ast.BinaryOp{Left: expr, Right: nil}

	}

	return expr.Left

}

// factor : ('+' | '-' | NOT ) factor | power
func (p *Parser) factor() ast.Noder {

	// p.expectedToken([]string{token.PLUS, token.MINUS, token.K_NOT, token.INT}, p.syntaxError("syntax error '"+p.currentToken.Value+"'"))

	if p.currentToken.Type == token.PLUS ||
		p.currentToken.Type == token.MINUS ||
		p.currentToken.Type == token.K_NOT {
		expr := &ast.UnaryOp{}
		expr.Operator = p.nextToken().Type
		expr.Value = p.factor()
		return expr
	} else {
		return p.power()
	}

}

// power : atom ('**' factor)?
func (p *Parser) power() ast.Noder {

	atom := p.atom()

	if p.currentToken.Type == token.POWER {
		expr := &ast.BinaryOp{}
		expr.Left = atom
		expr.Operator = p.nextToken().Value
		expr.Right = p.factor()
		return expr
	}

	return atom

}

// atomSimple : Integer | Float | String
func (p *Parser) atomSimple() ast.Noder {
	atom := &ast.Node{Type: p.currentToken.Type, Value: p.currentToken.Value}
	p.nextToken()
	return atom
}

// atomBool : Bool
func (p *Parser) atomBool() ast.Noder {
	atom := &ast.Node{Type: token.BOOL, Value: p.currentToken.Value}
	p.nextToken()
	return atom
}

// atomParens : '(' expression ')'
func (p *Parser) atomParens() ast.Noder {
	// Parsing expression surrounded by parentheses
	p.expectedToken([]string{token.LPAREN}, p.syntaxError("invalid syntax, expecting '('"))
	p.nextToken() // Move past first paren

	expr := p.expression()

	p.expectedToken([]string{token.RPAREN}, p.syntaxError("invalid syntax, expecting ')'"))
	p.nextToken()

	return expr
}

// identifierList : Identifier ( ',' Identifier )*
func (p *Parser) identifierList() ast.Noder {

	list := &ast.List{}
	// Identifier
	if p.currentToken.Type == token.IDENT {

		list.Values = append(list.Values, p.atomIdentifier())

	}

	// ( ',' Identifer )*
	for p.currentToken.Type == token.COMMA {
		p.nextToken()
		list.Values = append(list.Values, p.atomIdentifier())
	}

	return list

}

// atomIdentifier : Identifier
//                | Identifier ( '(' expressionList? ')' )+
//                | Identifier ( '[' expression ']' )+
func (p *Parser) atomIdentifier() ast.Noder {

	if p.peek().Type == token.LPAREN {

		atom := &ast.BinaryOp{}
		atom.Left = &ast.Node{Type: token.IDENT, Value: p.currentToken.Value}
		// Identifier
		p.nextToken()
		for p.currentToken.Type == token.LPAREN {
			atom.Operator = "()"
			// (
			p.nextToken()
			// expressionList?
			if p.currentToken.Type != token.RPAREN {
				atom.Right = p.expressionList()

				p.expectedToken([]string{token.RPAREN}, p.syntaxError("invalid syntax, expecting ')'"))
				p.nextToken()
			} else {
				p.nextToken()
			}
			atom = &ast.BinaryOp{Left: atom}
		}

		return atom.Left
	} else if p.peek().Type == token.LBRACK {

		atom := &ast.BinaryOp{}
		atom.Left = &ast.Node{Type: p.currentToken.Type, Value: p.currentToken.Value}
		// Identifier
		p.nextToken()
		for p.currentToken.Type == token.LBRACK {
			atom.Operator = "[]"
			// [
			p.nextToken()
			// expression
			if p.currentToken.Type != token.RBRACK {
				atom.Right = p.expression()

				p.expectedToken([]string{token.RBRACK}, p.syntaxError("invalid syntax, expecting ']'"))
				p.nextToken()
			} else {
				p.unexpectedToken([]string{token.RBRACK}, p.syntaxError("invalid syntax, unexpected ']'"))
				p.nextToken()
			}
			atom = &ast.BinaryOp{Left: atom}
		}

		return atom.Left
	}

	p.expectedToken([]string{token.IDENT}, p.syntaxError("expected identifier after "+p.previousToken.Value))
	atom := &ast.Node{Type: p.currentToken.Type, Value: p.currentToken.Value}
	p.nextToken()

	return atom

}

// atomAnonFunc : Func '(' parameterList? ')' block ( '(' expressionList? ')' )+
func (p *Parser) atomAnonFunc() ast.Noder {

	p.nextToken()

	function := &ast.Function{}

	function.Name = "@func" + strconv.FormatInt(int64(p.anonFuncCount), 10)
	p.anonFuncCount++

	p.expectedToken([]string{token.LPAREN}, p.syntaxError("expected ("))
	p.nextToken()

	// parameterList?
	if p.currentToken.Type != token.RPAREN {
		function.Parameters = *(p.identifierList()).(*ast.List)
	}

	p.expectedToken([]string{token.RPAREN}, p.syntaxError("expected )"))
	p.nextToken()

	function.Body = p.block()

	if p.currentToken.Type == token.LPAREN {
		atom := &ast.BinaryOp{}
		atom.Left = function

		for p.currentToken.Type == token.LPAREN {
			atom.Operator = "()"
			// (
			p.nextToken()
			// expressionList?
			if p.currentToken.Type != token.RPAREN {
				atom.Right = p.expressionList()

				p.expectedToken([]string{token.RPAREN}, p.syntaxError("invalid syntax, expecting ')'"))
				p.nextToken()
			} else {
				p.nextToken()
			}
			atom = &ast.BinaryOp{Left: atom}
		}
		return atom.Left
	}

	return function

}

// atomList : '[' (expression ( ',' expression )*)? ']'
func (p *Parser) atomList() ast.Noder {

	list := &ast.List{}

	p.nextToken()

	if p.currentToken.Type != token.RBRACK {

		list.Values = append(list.Values, p.expression())

		if p.currentToken.Type == token.RBRACK {
			p.nextToken()
		}
		// ( ',' expression )*
		for p.currentToken.Type == token.COMMA {
			p.nextToken()
			list.Values = append(list.Values, p.expression())
		}

		p.expectedToken([]string{token.RBRACK}, p.syntaxError("expecting ']'"))

	}
	return list

}

// atom : atomSimple | atomBool | atomParens | atomIdentifier | atomAnonFunc | atomList
func (p *Parser) atom() ast.Noder {

	p.expectedToken([]string{
		token.IDENT, token.INT, token.STRING, token.FLOAT, // type tokens
		token.K_TRUE, token.K_FALSE, token.K_FUNC, token.K_INT, // keyword tokens
		token.LPAREN, token.RBRACE, token.LBRACK}, p.syntaxError("found '"+p.currentToken.Value+"'"))

	if p.currentToken.Type == token.STRING ||
		p.currentToken.Type == token.FLOAT ||
		p.currentToken.Type == token.INT {
		return p.atomSimple()

	} else if p.currentToken.Type == token.K_TRUE ||
		p.currentToken.Type == token.K_FALSE {

		return p.atomBool()

	} else if p.currentToken.Type == token.IDENT ||
		p.currentToken.Type == token.K_INT {
		return p.atomIdentifier()

	} else if p.currentToken.Type == token.K_FUNC {
		return p.atomAnonFunc()

	} else if p.currentToken.Type == token.LPAREN {

		return p.atomParens()

	} else if p.currentToken.Type == token.LBRACK {

		return p.atomList()

	}

	return &ast.Node{Type: "NIL", Value: "nil"}

}
